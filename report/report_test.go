package report;

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/gorilla/mux"	
	"io/ioutil"
	"crypto/tls"
	"time"	
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strconv"
	"bytes"	
)

func testJsonReply(t *testing.T, r *http.Response, st int, cnt interface{}) {
	assert.Equal(t, st, r.StatusCode, "Expected status " + strconv.Itoa(st) + " got " + strconv.Itoa(r.StatusCode))
	bytes, err := ioutil.ReadAll(r.Body)
	assert.Nil(t, err, "Unable to read the body")
	json.Unmarshal(bytes, cnt)
}

func postJson(t *testing.T, client *http.Client, url string, cnt interface{}) *http.Response {		
		buf, err := json.Marshal(cnt)
		assert.Nil(t, err, "Unable to encode the body")
		resp, err := client.Post(url, "application/json", bytes.NewReader(buf))
		assert.Nil(t, err, "Unable to post the request")
		return resp
}

func TestWorkflow(t *testing.T) {
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
	rs := NewMock()
	//2 reports for toto
	m1,_ := rs.New("midterm","toto", time.Now())
	m2,_ := rs.New("final","toto", time.Now())	
	r := mux.NewRouter()
	srv := httptest.NewTLSServer(r)
	defer srv.Close()
	RegisterHandlers(rs, nil, r)			
	client := &http.Client{Transport: tr}
	
	resp, _ := client.Get(srv.URL + "/reports/tito")	
	assert.Equal(t, resp.StatusCode, 404,"")

	resp, _ = client.Get(srv.URL + "/reports/toto")	
	assert.Equal(t, resp.StatusCode, 200,"")
	
	//List
	var meta []MetaData
	testJsonReply(t, resp, 200, &meta)	
	assert.Equal(t, len(meta), 2, "")
	for _, m := range(meta) {
		if (m.Kind == "midterm") {
			assert.Equal(t, m, m1)
		} else if (m.Kind == "final") {
			assert.Equal(t, m, m2)
		} else {
			t.Fatalf("Unexpected: %s\n", m)
		}
	}

	//try to grade m1, no way because not posted
	resp = postJson(t, client, srv.URL + "/reports/toto/midterm/mark",5)
	assert.Equal(t, resp.StatusCode, 403, "")
	
	//Unknown student	
	resp = postJson(t, client, srv.URL + "/reports/titi/midterm/mark", 5)
	assert.Equal(t, resp.StatusCode, 404,"")

	//Unknown report	
	resp = postJson(t, client, srv.URL + "/reports/toto/foo/mark", 5)
	assert.Equal(t, resp.StatusCode, 404,"")

	//Invalid grade	
	resp = postJson(t, client, srv.URL + "/reports/toto/midterm/mark", -5)
	assert.Equal(t, resp.StatusCode, 403,"")

	//delay
	now := time.Now()
	resp = postJson(t, client, srv.URL + "/reports/toto/midterm/deadline", now)
	assert.Equal(t, resp.StatusCode, 200)
	resp, _ = client.Get(srv.URL + "/reports/toto")	
	testJsonReply(t, resp, 200, &meta)		
	for _, m := range(meta) {
		if (m.Kind == "midterm") {			
			assert.Equal(t, m.Deadline, now)
		} 
	}

	//Let's push content
	resp, _= client.Post(srv.URL + "/reports/toto/foo/content", "application/pdf", bytes.NewReader([]byte{0}))
	assert.Equal(t, 404, resp.StatusCode, "unknown report")
	resp, _ = client.Post(srv.URL + "/reports/titi/foo/content", "application/pdf", bytes.NewReader([]byte{0}))
	assert.Equal(t, 404, resp.StatusCode, "unknown student")
	resp, _ = client.Post(srv.URL + "/reports/titi/foo/content", "application/txt", bytes.NewReader([]byte{0}))
	assert.Equal(t, 415, resp.StatusCode, "unsupported format")

	resp, _= client.Post(srv.URL + "/reports/toto/midterm/content", "application/pdf", bytes.NewReader([]byte{1,2}))
	assert.Equal(t, 200, resp.StatusCode, "")	
	//resp, _ = client.Get(srv.URL + "/reports/toto/midterm/content")
	//b,_ := ioutil.ReadAll(resp.Body)	
	//assert.Equal(t, b, []byte{1,2})

	//grade		
	resp = postJson(t, client, srv.URL + "/reports/toto/midterm/mark", 5)
	assert.Equal(t, 200, resp.StatusCode,"")
	resp, _ = client.Get(srv.URL + "/reports/toto")	

	testJsonReply(t, resp, 200, &meta)	
	for _, m := range(meta) {
		if (m.Kind == "midterm") {			
			assert.Equal(t, m.Grade, 5)
		} 
	}
}



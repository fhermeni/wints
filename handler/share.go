package handler;

import (
	"net/http"
	"io/ioutil"
	"log"
)

func fileReply(w http.ResponseWriter, mime, filename string, cnt []byte) {
	w.Header().Set("Content-type", mime)
	w.Header().Set("Content-disposition", "attachment; filename=" + filename)
	_, err := w.Write(cnt)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		log.Println("Unable to send file '" + filename + "': " + err.Error())
	}	
}

func requiredContent(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	cnt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, err
	}
	if len(cnt) == 0 {
		http.Error(w, "", http.StatusBadRequest)
		return []byte{}, err
	}
	return cnt, nil
}
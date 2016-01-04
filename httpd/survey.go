package httpd

import (
	"html/template"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/fhermeni/wints/schema"
)

func (ed *HTTPd) survey(w http.ResponseWriter, r *http.Request) {
	//Cache ?
	kind := r.URL.Query().Get("kind")
	s := schema.Survey{}
	if _, err := toml.DecodeFile("static/surveys/"+kind+".toml", &s); err != nil {
		ed.not.Log.Log("root", "Unable to read the survey questions", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	tpl, err := template.New("survey.html").ParseFiles("static/surveys/survey.html")
	if err != nil {
		ed.not.Log.Log("root", "Unable to read the survey template", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	if err := tpl.Execute(w, s); err != nil {
		ed.not.Log.Log("root", "Unable to run the template", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

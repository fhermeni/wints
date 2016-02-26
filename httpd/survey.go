package httpd

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/fhermeni/wints/logger"
	"github.com/fhermeni/wints/schema"
)

func (ed *HTTPd) survey(w http.ResponseWriter, r *http.Request) {
	//Cache ?
	kind := r.URL.Query().Get("kind")
	s := schema.Survey{}
	if _, err := toml.DecodeFile("assets/surveys/"+kind+".toml", &s); err != nil {
		logger.Log("event", "survey", "Unable to read the questions for '"+kind+"'", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	tpl, err := template.New("survey.html").Funcs(
		template.FuncMap{
			"html": func(value interface{}) template.HTML {
				return template.HTML(fmt.Sprint(value))
			},
		}).ParseFiles("assets/surveys/survey.html")

	if err != nil {
		logger.Log("event", "survey", "Unable to read the template for '"+kind+"'", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	if err := tpl.Execute(w, s); err != nil {
		logger.Log("event", "survey", "Unable to run the template for '"+kind+"'", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

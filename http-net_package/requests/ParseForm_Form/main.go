package main

import (
	"html/template"
	"log"
	"net/http"
)

type rqt int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func (m rqt) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func main() {

	var req rqt
	http.ListenAndServe(":8080", req)
}

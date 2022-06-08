package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

//creating a type with underlying type int; used create a type Handler (interface)
type reqv int

func (m reqv) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}

	// LOOK AT THE DOCUMENTATION FOR http.Request
	data := struct {
		Method        string   //METHOD -> look at the docs for http.REQUEST
		URL           *url.URL //URL -> look at the docs for http.REQUEST
		Submissions   url.Values
		Header        http.Header //returns a map[string]string
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d reqv
	http.ListenAndServe(":8080", d)
}

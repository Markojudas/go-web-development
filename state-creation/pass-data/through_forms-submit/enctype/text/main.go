package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

/* type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
} */

func foo(w http.ResponseWriter, req *http.Request) {

	//BODY
	byteslice := make([]byte, req.ContentLength) //the since of the slice is the content length

	req.Body.Read(byteslice) //reads INTO byteslice
	body := string(byteslice)

	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

/*
	output:

	BODY: first=Jose last=Hernandez subscribe=on

	**NOTE: used for debugging never in production
*/

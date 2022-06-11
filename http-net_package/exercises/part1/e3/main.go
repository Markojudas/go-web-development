package main

import (
	"html/template"
	"log"
	"net/http"
)

func parseTemplate() (*template.Template, error) {
	tpl, err := template.ParseFiles("index.gohtml")

	return tpl, err
}

func execTemplate(tpl *template.Template, w http.ResponseWriter, str string) error {

	err := tpl.ExecuteTemplate(w, "index.gohtml", str)

	return err
}

func errorHandler(err error) {
	if err != nil {
		log.Fatalln("error:", err)
	}
}

// HomePage
func homePage(w http.ResponseWriter, req *http.Request) {

	tpl, err := parseTemplate()
	errorHandler(err)

	err = execTemplate(tpl, w, "Welcome to my site!")
	errorHandler(err)
}

// Dog/
func dogHandler(w http.ResponseWriter, req *http.Request) {
	tpl, err := parseTemplate()
	errorHandler(err)

	err = execTemplate(tpl, w, "woof woof!")
	errorHandler(err)
}

// me/
func meHandler(w http.ResponseWriter, req *http.Request) {

	tpl, err := parseTemplate()
	errorHandler(err)

	err = execTemplate(tpl, w, "Markojudas")
	errorHandler(err)
}

func main() {

	// here I am converting the functions homePage, dogHandler, & meHandler to type http.HandlerFunc(ResponsiveWriter, *Request)
	// type http.HandlerFunc(ResponsiveWriter, *Request) implements http.Handler interface since it has the method ServeHTTP(ResponsiveWriter, *Request)
	// http.Handle functions takes a pattern (Type string, route!) and a http.Handler.
	// since http.HandlerFunc is also type http.Handler the below code works!
	http.Handle("/", http.HandlerFunc(homePage))
	http.Handle("/dog/", http.HandlerFunc(dogHandler))
	http.Handle("/me/", http.HandlerFunc(meHandler))

	http.ListenAndServe(":8080", nil)
}

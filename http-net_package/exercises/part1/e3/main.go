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

	http.Handle("/", http.HandlerFunc(homePage))
	http.Handle("/dog/", http.HandlerFunc(dogHandler))
	http.Handle("/me/", http.HandlerFunc(meHandler))

	http.ListenAndServe(":8080", nil)
}

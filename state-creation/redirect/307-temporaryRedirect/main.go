package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("your request method at bar: ", req.Method)

	//Process form data
	// w.Header().Set("Location", "/")
	// w.WriteHeader(http.StatusSeeOther)

	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your response method at barred: ", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)

	//the form action sends the form to /bar - METHOD wil be POST
	//at /bar we set the location to "/"" and we're using http.StatusTemporaryRedirect the method remains POST}
}

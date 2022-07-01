package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	io.WriteString(w, req.Method)
}

func foo(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "\".\"") -> unecessary
	str := `Here is some text from foo`
	fmt.Fprintln(w, str)
	fmt.Println(req.Header)
}

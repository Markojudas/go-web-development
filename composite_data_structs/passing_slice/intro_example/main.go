package main

/*

	to pass a composite type (ie slince):
		range
		use:
		{{range .}}
		<li>{{.}}</li>
		{{end}}
*/

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	sages := []string{
		"Gandhi",
		"MLK",
		"Buddha",
		"Jesus",
		"Muhammad",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

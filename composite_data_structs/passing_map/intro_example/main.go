package main

/*

	to pass a composite type (ie map):
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

	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

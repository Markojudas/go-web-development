package main

/*

	to pass a composite type (ie slice):
		range
		use:
		{{range $index, $element := .}}
		<li>{{$index - $element}}</li>
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

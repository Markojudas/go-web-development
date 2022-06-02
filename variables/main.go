package main

/*

	Variables in template:

	assign the data to the variable: {{$variable := .}}
	use the variable {{$variable}}
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

	str := "Release self-focus; embrace other-focus"

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", str)
	error_check(err)

	// create html file
	fp, err := os.Create("index.html")
	error_check(err)
	defer fp.Close()

	err = tpl.ExecuteTemplate(fp, "tpl.gohtml", str)
	error_check(err)
}

func error_check(err error) {

	if err != nil {
		log.Fatalln(err)
	}
}

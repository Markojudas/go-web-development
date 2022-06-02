package main

import (
	"log"
	"os"
	"text/template"
)

/*
	to pass data: the place holder on the template: {{.}}
*/

var tpl *template.Template //package scope for init

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	//the third argument is the data being passed
	//can only pass 1 data -> can be an agregate data type (struct, list...)
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatal(err)
	}
}

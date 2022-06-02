package main

/*

	to pass a composite type (ie a slice of struct): with variables
		use:

		range over the slice
		and access data with the corrspeonding fields

		<ul>
    		{{range .}}
    		<li>{{.Name}} - {{.Motto}}</li>
    		{{end}}
		<ul>

*/

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "The Belief of no Beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the Change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil",
	}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

package main

/*

	to pass a composite type (ie struct): with variables
		use:
		<ul>
			{{$x := .Name}}
			{{$y := .Motto}}
			<li>{{$x}} - {{$y}}</li>
		</ul>

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

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}

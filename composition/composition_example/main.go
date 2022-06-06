package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	y1 := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"COP4610", "Operating Systems Principles", "3"},
				{"COP4550", "Principles of Programming Language", "3"},
				{"COP4375", "Net-Centric Computing", "3"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"HEL6660", "It's Hot as hell", "3"},
				{"HEL5550", "I am tired of Claims", "3"},
				{"HEL7770", "I never falter", "3"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y1)
	if err != nil {
		log.Fatalln(err)
	}

}

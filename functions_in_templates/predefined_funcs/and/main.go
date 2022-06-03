package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type user struct {
	Name  string
	Motto string
	Admin bool
}

func main() {

	u1 := user{
		Name:  "Jose",
		Motto: "I know my truth",
		Admin: true,
	}

	u2 := user{
		Name:  "Lucifer",
		Motto: "Be true to thy self",
		Admin: false,
	}

	u3 := user{
		Name:  "",
		Motto: "Nothing",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}

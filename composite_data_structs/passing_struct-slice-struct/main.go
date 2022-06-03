package main

/*

	to pass a composite type (ie a struct with a slice of struct): with variables
		use:

		you're passing the struct
		...access the field, which is a slice of struct. so range that field
		then access the correspoding field

	<ul>
    	{{range .Wisdom}}
    	<li>{{.Name}} - {{.Motto}}</li>
    	{{end}}
	<ul>

	<ul>
    	{{range .Transport}}
    	<li>{{.Manufacturer}} - {{.Model}} - {{.Doors}} door</li>
    	{{end}}
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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
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

	veh1 := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	veh2 := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}
	cars := []car{veh1, veh2}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	/*
		Alternative way!
		let's say that the item struct is not declared on top with package scope

		data := struct{
			Wisdom []sage
			Transport []car
		}{
			sages,
			cars,
		}

	*/

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

	//  just creating an html file! :)

	fp, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer fp.Close()

	err = tpl.ExecuteTemplate(fp, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

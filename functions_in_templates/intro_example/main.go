package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

var tpl *template.Template

//Using template.FuncMap{} function from package text/template
//key is a string and value could be anything(function
//the key will be the name given to the function so we can access it on the template
var func_map = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

//First three takes a string and returns the first 3 characters
func firstThree(s string) string {
	s = strings.TrimSpace(s) //returns a slice composed of the string removing all white spaces

	//making sure the string, without whitespace has 3 characters or more
	if len(s) >= 3 {
		s = s[:3]
	}

	return s
}

func init() {
	//using template.New.Funcs
	// the functions NEED to be present/attach before parsing
	// .Must gives a *template.Template
	// but to pass the function to the template we need to call the .Funcs function.
	// to access that func we need a type *template.Template
	// that's why we also use template.New() that reutnrs a *template.Template
	tpl = template.Must(template.New("").Funcs(func_map).ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "It really doesn't matter what I put here",
	}

	satan := sage{
		Name:  "Lucifer",
		Motto: "Be true to thy self",
	}

	wolf := sage{
		Name:  "Big Bad Wolf",
		Motto: "Huffing and Puffing, I'll blow your house in",
	}

	veh1 := car{
		Manufacturer: "Scion",
		Model:        "FR-S",
		Doors:        2,
	}

	veh2 := car{
		Manufacturer: "Toyota",
		Model:        "Yaris",
		Doors:        4,
	}

	sages := []sage{buddha, gandhi, satan, wolf}
	cars := []car{veh1, veh2}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}

package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

//this will run first when the program starts
func init() {
	//template.Must wraps a call to a function that returns (*Template, error) it panics if err != nil; it does error checking
	//use for variable initialization
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Error executing:", err)
	}

	//since tpl is a "cointainer" with multiple files
	//need to specify which template to use

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatal("Error excecuting:", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatal("Error executing:", err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("error executing:", err) // the first one on the container
	}
}

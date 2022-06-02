package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// To parse a group templates at once. "templates/*" selects everything from the directory
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatal("Error parsing file:", err)
	}

	err = tpl.Execute(os.Stdout, nil)
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

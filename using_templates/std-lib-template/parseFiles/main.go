package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseFiles("one.gohtml")
	if err != nil {
		log.Fatal("Error parsing file:", err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Error executing:", err)
	}

	//parings more than one file!

	tpl, err = tpl.ParseFiles("two.gohtml", "three.gohtml")
	if err != nil {
		log.Fatal("Error parsing file(s)", err)
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

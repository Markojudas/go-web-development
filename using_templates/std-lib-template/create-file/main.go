package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseFiles("tpl.gohtml") //returns *Template, error

	if err != nil {
		log.Fatal("error opening file:", err)
	}

	fp, err := os.Create("index.html")

	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer fp.Close()

	err = tpl.Execute(fp, nil)
	if err != nil {
		log.Fatal("error executing the template:", err)
	}
	fmt.Println()
}

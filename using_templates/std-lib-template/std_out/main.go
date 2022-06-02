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

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("error executing the template to", os.Stdout, ":", err)
	}
	fmt.Println()
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	name := os.Args[1]
	fmt.Println(os.Args[0]) //name of running process
	fmt.Println(os.Args[1]) //argument passed

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World</title>
	</head>
	<body>
	<h1> ` + name + `</h1>
	</body>
	</html>
	`
	fp, err := os.Create("index.html")

	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer fp.Close() //defer the closing of file for the end

	io.Copy(fp, strings.NewReader(tpl))

	//at terminal:
	//go run main.go <name>
}

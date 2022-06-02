package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	name := "Jose Hernandez"

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
}

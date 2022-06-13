package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	// the String to reader
	str := `My name is Lucider,
please take my hand
`

	// creating a reader (pointer to it)
	readerPtr := strings.NewReader(str)

	//io.Copy
	_, err := io.Copy(os.Stdout, readerPtr)
	if err != nil {
		log.Fatal(err)
	}
}

/*
	output:
	My name is Lucider,
	please take my hand
*/

package main

import (
	"io"
	"log"
	"os"
)

func main() {
	_, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}

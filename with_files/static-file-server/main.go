package main

import (
	"log"
	"net/http"
)

func main() {

	// http.ListenAndServe returns an error! we need to catch it!
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}

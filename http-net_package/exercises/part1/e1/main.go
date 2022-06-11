package main

import (
	"io"
	"net/http"
)

// home page func
func homePage(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the HomePage!")
}

// dog/
func dogHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am the dog, dawg!")
}

// me/
func meHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Markojudas")
}

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/me/", meHandler)

	http.ListenAndServe(":8080", nil)
}

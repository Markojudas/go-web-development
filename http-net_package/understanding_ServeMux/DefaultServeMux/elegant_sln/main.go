package main

import (
	"io"
	"net/http"
)

/*
MUCH ELEGANT SOLUTION:
Using HandleFunc and the DefaultServeMux aka, passing nil to http.ListenAndServe()

creating functions with the params ResponseWriter and *Request
Using the http.HandleFunc() method
Read Docs!
*/

func dogHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am the DOG!")
}

func catHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am the CAT!")
}

func main() {

	http.HandleFunc("/dog/", dogHandler) //any subpath will work
	http.HandleFunc("/cat", catHandler)  //any subpath will NOT work, does not have a trailing "/"

	http.ListenAndServe(":8080", nil)
}

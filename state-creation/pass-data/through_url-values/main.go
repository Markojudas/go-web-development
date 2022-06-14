package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	val := req.FormValue("q") //retriving the value for the URL
	io.WriteString(w, "d my search: "+val)
}

//visit:
// localhost:8080/?q=dog

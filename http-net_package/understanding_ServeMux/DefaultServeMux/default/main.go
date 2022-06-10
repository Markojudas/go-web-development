package main

import (
	"io"
	"net/http"
)

/*
	Using the default mux : passing nil to method!
*/

type hotdog int
type hotcat int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am the DOG!")
}

func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am the CAT!")
}

func main() {

	var d hotdog
	var c hotcat

	//NOT CREATING THE MUX!!! BUT USING THE DEFAULT!

	http.Handle("/dog/", d) //any subpath will work
	http.Handle("/cat", c)  //any subpath will NOT work, does not have a trailing "/"

	http.ListenAndServe(":8080", nil)
}

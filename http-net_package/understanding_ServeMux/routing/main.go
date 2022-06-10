package main

/*
NOT THE BEST SOLUTION

*/

import (
	"io"
	"net/http"
)

// create a type of underlying type int
type hotdog int

//attached ServeHTTP(responseWriter, *Request) to the type to implement the Handler interface
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	//example implementation for routing
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy doggy doggy") //writing the string to the responseWriter
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}
}

func main() {
	//creating variable of type hotdog that implements Handler interface
	var d hotdog
	http.ListenAndServe(":8080", d)
}

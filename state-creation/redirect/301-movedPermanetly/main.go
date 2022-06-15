package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("your request method at bar: ", req.Method)

	http.Redirect(w, req, "/", http.StatusMovedPermanently) //same thing!
}

/*
	output:

	Your request method at foo: GET

	your request method at bar:  GET
	Your request method at foo: GET

	Your request method at foo: GET

	the first output is when we went on the browser to: localhost:8080/
	the second and third is because when we went to localhost:8080/bar it REDIRECTED to localhost:8080/

	the last output is because it was MOVED PERMANENTLY to "/", even though we went to "/bar"... the browser cached the move

*/

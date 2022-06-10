package main

/*
Better Solution but could be better!

*/

import (
	"io"
	"net/http"
)

//HANDLERS
//creating type with arbitrary name with underlying type int
type hotdog int

//attaching ServeHTTP method to the type so it implements the Handler interface
func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

//creating yet another type of arbitrary name with underlying type int
type hotcat int

//attaching ServeHTTP method to the type so it implements the Handler interface
func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat cat cat")
}

//END OF HANDLERS

func main() {

	//creating variables of both types
	var d hotdog
	var c hotcat

	//creating a mux (mutliplexer) for routing, using the NewServeMux() function
	//returns a *ServeMux
	mux := http.NewServeMux()

	//passing the variables, that implements the Handler interface, to the mux.Handle() method
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	// passing the mux, with the registered handlers and their patterns
	http.ListenAndServe(":8080", mux)

	/*
		/dog/ -> will work dog/some/thing/else
		whereas
		/cat -> will not work  cat/some/thing/else ::: 404 page not found

		Reason: trailing "/"

	*/

}

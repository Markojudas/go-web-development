package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// now TYPE hotdog is also a HANDLER!!!!!
func (m hotdog) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(wr, "Any code you want!")
}

func main() {
	var d hotdog

	//(addr string, handler)
	http.ListenAndServe(":8080", d)
}

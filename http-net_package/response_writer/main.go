package main

/*
Read the docs of http.ResponseWriter
Read the docs of http.ListenAndServe
*/

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	//setting the headers
	w.Header().Set("markojudas-key", "Finally I am king of the world")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any Code you want in this func</h1>") //writes to the body
}

func main() {
	var d hotdog                    //Handler!
	http.ListenAndServe(":8080", d) //passing handler ListenAndServe
}

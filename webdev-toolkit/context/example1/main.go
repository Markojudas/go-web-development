package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

/*
output:

context.Background.WithValue(type *http.contextKey, val <not Stringer>).WithValue(type *http.contextKey, val 127.0.0.1:8080).WithCancel.WithCancel

*/

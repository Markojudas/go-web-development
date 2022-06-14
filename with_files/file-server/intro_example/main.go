package main

import (
	"fmt"
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="Rufus.jpg">`)
	fmt.Println("METHOD:", req.Method)
	fmt.Println("URL:", req.URL)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil)
}

/*
	Output:
	localhost:8080 -> shows the content of the directory
	localhost:8080/dog -> serves the img written to responsewriter

*/

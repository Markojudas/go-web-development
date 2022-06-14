package main

import (
	"fmt"
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/Rufus.jpg">`)
	fmt.Println("METHOD:", req.Method)
	fmt.Println("URL:", req.URL)
}

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assests"))))
	http.ListenAndServe(":8080", nil)
}

/*
	localhost:8080/resources/
	will strop away the prefix /resources (/resources/Rufus.jpg)
	and look on ./assets

*/

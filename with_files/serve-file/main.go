package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/Rufus.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {

	//without having to open file!!
	http.ServeFile(w, req, "Rufus.jpg")
}

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/Rufus.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

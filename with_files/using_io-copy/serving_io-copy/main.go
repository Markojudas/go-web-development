package main

import (
	"io"
	"net/http"
	"os"
)

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="Rufus.jpg">
	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	fp, err := os.Open("Rufus.jpg") //opening file, file ptr -> a file IS a reader!
	if err != nil {
		http.Error(w, "file not found", 404)
	}
	defer fp.Close()

	io.Copy(w, fp) //a responseWriter IS a writer so this works!
}

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/Rufus.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

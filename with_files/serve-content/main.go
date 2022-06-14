package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/Rufus.jpg">`)
	reqURL := fmt.Sprintln("URL:", req.URL)
	reqMETHOD := fmt.Sprintln("Method:", req.Method)
	io.WriteString(w, reqURL)
	io.WriteString(w, "\n")
	io.WriteString(w, reqMETHOD)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	//Open and serve

	//Open
	fd, err := os.Open("Rufus.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer fd.Close()

	// File info
	finfo, err := fd.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, req, finfo.Name(), finfo.ModTime(), fd)
}

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/Rufus.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

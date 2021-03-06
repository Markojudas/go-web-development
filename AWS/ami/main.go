package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from AWS.")
	io.WriteString(w, "\n")
	io.WriteString(w, req.Host)
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
	io.WriteString(w, "\n")
	io.WriteString(w, req.Host)
}

func instance(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		fmt.Println(err)
		return
	}

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()
	io.WriteString(w, string(bs))
	io.WriteString(w, "\n")
	io.WriteString(w, req.Host)
}

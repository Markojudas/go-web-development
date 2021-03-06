package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("this is an example server.\n"))
	io.WriteString(w, req.Method)
}

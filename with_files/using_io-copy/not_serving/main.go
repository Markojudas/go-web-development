package main

import (
	"io"
	"net/http"
)

//creating a TYPE func(ResponseWriter, *Request)
func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/hmtl; charset=UTF-8")

	io.WriteString(w, `
	<!--Not Serving from our server-->
	<img src="https://ifrit.markojudas.com/resources/BS.png">
	`)
}

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {
	val := req.FormValue("q")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="post">
		<input type="text" name="q">
		<input type="submit">
	</form>
	<br>`+val)
}

// POST -> data sent through the body (url stays the same)
// GET -> data sent through the URL (url gets appended with /?q=<value>)

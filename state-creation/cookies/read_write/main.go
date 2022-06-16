package main

import (
	"fmt"
	"net/http"
)

func main() {

	// first write the cookie
	http.HandleFunc("/", set)

	// read written cookie
	http.HandleFunc("/read", read)

	http.Handle("/favion.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

// setting a cookie
func set(w http.ResponseWriter, req *http.Request) {
	//composite literal to pass a cookie to SetCookie method; only setting the name and value
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})

	fmt.Fprintln(w, "Cookie Written - Check browser")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
	fmt.Fprintln(w, "METHOD:", req.Method, "\n", "URL:", req.URL)
}

// read the cookie!
func read(w http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintln(w, "YOUR COOKIE:", cookie)
}

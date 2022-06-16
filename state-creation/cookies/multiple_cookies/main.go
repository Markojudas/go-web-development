package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// first write the cookie
	http.HandleFunc("/", set)

	// read written cookie
	http.HandleFunc("/read", read)

	// setting multiple cookies
	http.HandleFunc("/multiCookies", multiCookies)

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
	fmt.Fprintln(w, "METHOD:", req.Method, "\n", "URL:", req.URL)
}

// read the cookie!
func read(w http.ResponseWriter, req *http.Request) {

	cookie1, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1", cookie1)
	}

	cookie2, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2", cookie2)
	}

	cookie3, err := req.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #3", cookie3)
	}
}

func multiCookies(w http.ResponseWriter, req *http.Request) {
	// setting two new cookies at once

	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "some other value about general things",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "some other value about specific things",
	})

	fmt.Fprintln(w, "COOKIES WRITTEN - CHECK BROWSER")
	fmt.Fprintln(w, "METHOD:", req.Method, "\n", "URL:", req.URL)

}

package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

/*
	Creating a counter to keep track the amount a user visits
*/

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler()) // needed because if not handled it will be caught by "/"
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	//requesting Cookie
	cookie, err := req.Cookie("my-cookie")

	//if not found
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}

	// Converting the cookie value into an INT
	count, err := strconv.Atoi(cookie.Value)

	if err != nil {
		log.Fatal(err)
	}

	count++
	// converting the cookie back to string
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)

	io.WriteString(w, cookie.Value)
}

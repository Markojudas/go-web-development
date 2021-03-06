package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewV4() //creating a session ID, unique session ID
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//Secure : true
			HttpOnly: true, // no access with javascript
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}

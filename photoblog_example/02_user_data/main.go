package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	cookie := getCookie(w, req)

	// passing the value of the cookie as the DATA to the template
	tpl.ExecuteTemplate(w, "index.gohtml", cookie.Value)
}

// helper function to set the cookie and returns it
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {

	cookie, err := req.Cookie("session")
	if err != nil {
		sessionID := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "sessions",
			Value: sessionID.String(),
		}
		http.SetCookie(w, cookie)
	}

	return cookie
}

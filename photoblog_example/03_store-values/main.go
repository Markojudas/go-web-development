package main

import (
	"html/template"
	"net/http"
	"strings"

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

	//using helper function to get the cookie
	cookie := getCookie(w, req)

	//using helper function to STORE value
	cookie = appendValue(w, cookie)

	//getting the values by spliting the cookie.Value
	xs := strings.Split(cookie.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)

}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {

	cookie, err := req.Cookie("session")

	//if the cookie is not found err != nil
	if err != nil {

		// create the sessionID and cookie
		sessionID := uuid.NewV4()

		cookie = &http.Cookie{
			Name:  "sessions",
			Value: sessionID.String(),
		}

		http.SetCookie(w, cookie)
	}

	return cookie
}

func appendValue(w http.ResponseWriter, cookie *http.Cookie) *http.Cookie {

	// Values
	pic1 := "disneyland.jpg"
	pic2 := "atbeach.jpg"
	pic3 := "hollywood.jpg"

	//append the values
	str := cookie.Value
	// if the string does NOT contain the "value"
	if !strings.Contains(str, pic1) {
		str += "|" + pic1
	}
	if !strings.Contains(str, pic2) {
		str += "|" + pic2
	}

	if !strings.Contains(str, pic3) {
		str += "|" + pic3
	}

	cookie.Value = str

	http.SetCookie(w, cookie)

	return cookie
}

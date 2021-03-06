package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favincon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	//get cookie
	cookie, err := req.Cookie("session")

	//if cookie not found...
	if err != nil {
		//create cookie
		sID := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
	}

	//if the user exists already get user
	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}

	//process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		first := req.FormValue("firstname")
		last := req.FormValue("lastname")
		u = user{un, first, last}
		dbSessions[cookie.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {

	//get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

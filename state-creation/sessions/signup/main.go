package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// User Struct
type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

// for template
var tpl *template.Template

//creating the database // maps for the session
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

//parsing all templates
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	_user := getUser(req)
	tpl.ExecuteTemplate(w, "index.gohtml", _user)
}

func bar(w http.ResponseWriter, req *http.Request) {
	_user := getUser(req)

	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "bar.gohtml", _user)
}

func signup(w http.ResponseWriter, req *http.Request) {

	//check if already signed in
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//process form submissions
	if req.Method == http.MethodPost {

		//get form values
		_username := req.FormValue("username")
		_password := req.FormValue("password")
		_firstName := req.FormValue("firstname")
		_lastName := req.FormValue("lastname")

		// username taken?
		if _, ok := dbUsers[_username]; ok {
			http.Error(w, "username already taken", http.StatusForbidden)
			return
		}

		// cerate session
		sID := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = _username

		// store user in dbUsers
		_user := user{_username, _password, _firstName, _lastName}
		dbUsers[_username] = _user

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//only reachable when NOT logged IN and METHOD == GET
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

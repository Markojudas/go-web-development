package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User Struct
type user struct {
	UserName string
	Password []byte
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
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUsers["test@test.com"] = user{"test@test.com", bs, "James", "Bond"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
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
		//encrypt password first
		bs, err := bcrypt.GenerateFromPassword([]byte(_password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		_user := user{_username, bs, _firstName, _lastName}
		dbUsers[_username] = _user

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//only reachable when NOT logged IN and METHOD == GET
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	//process form submission
	if req.Method == http.MethodPost {
		_username := req.FormValue("username")
		_password := req.FormValue("password")

		//is there a username???
		_user, ok := dbUsers[_username]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		//does the entered passwoerd match the stored password?
		err := bcrypt.CompareHashAndPassword(_user.Password, []byte(_password))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		//create session
		sID := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = _username
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//METHOD == GET && not logged in
	tpl.ExecuteTemplate(w, "login.gohtml", nil)

}

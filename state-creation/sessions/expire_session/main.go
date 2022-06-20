package main

import (
	"html/template"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	username     string
	lastActivity time.Time
}

var tpl *template.Template

var dbUsers = map[string]user{}

var dbSessions = map[string]session{}

var dbSessionsCleaned time.Time

const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))

	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil) //chaning the port to 80 for AWS exercise, from 8080
}

func index(w http.ResponseWriter, req *http.Request) {
	_user := getUser(w, req)
	showSessions()
	tpl.ExecuteTemplate(w, "index.gohtml", _user)
}

func bar(w http.ResponseWriter, req *http.Request) {
	_user := getUser(w, req)

	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// this checks the user's Role attribute.
	if _user.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}

	showSessions()
	tpl.ExecuteTemplate(w, "bar.gohtml", _user)
}

func signup(w http.ResponseWriter, req *http.Request) {

	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var _user user

	if req.Method == http.MethodPost {
		_username := req.FormValue("username")
		_password := req.FormValue("password")
		_firstName := req.FormValue("firstname")
		_lastName := req.FormValue("lastname")
		_role := req.FormValue("role")

		// username taken?
		if _, ok := dbUsers[_username]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		cookie.MaxAge = sessionLength
		http.SetCookie(w, cookie)
		// dbSessions now also needs the time
		dbSessions[cookie.Value] = session{_username, time.Now()}

		//store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(_password), bcrypt.MinCost)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		_user = user{_username, bs, _firstName, _lastName, _role}
		dbUsers[_username] = _user

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "signup.gohtml", _user)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var _user user

	if req.Method == http.MethodPost {
		_username := req.FormValue("username")
		_password := req.FormValue("password")

		// is there a username?
		_user, ok := dbUsers[_username]

		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// does the entered password match the stored password
		err := bcrypt.CompareHashAndPassword(_user.Password, []byte(_password))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		cookie.MaxAge = sessionLength
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = session{_username, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "login.gohtml", _user)
}

func logout(w http.ResponseWriter, req *http.Request) {

	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	cookie, _ := req.Cookie("session")
	//delete the session
	delete(dbSessions, cookie.Value)

	//remove the cookie
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	// Clean up dbSessions
	if time.Since(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

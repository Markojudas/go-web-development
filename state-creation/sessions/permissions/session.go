package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) user {

	//get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(w, cookie)

	// if the user exists already, get user
	var _user user
	if _username, ok := dbSessions[cookie.Value]; ok {
		_user = dbUsers[_username]
	}
	return _user
}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}

	_username := dbSessions[cookie.Value]
	_, ok := dbUsers[_username]
	return ok
}

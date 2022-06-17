package main

import "net/http"

func getUser(req *http.Request) user {
	var _user user

	// get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		return _user
	}

	//if the user exists already, get user
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

package main

import (
	"fmt"
	"net/http"
	"time"

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
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)

	// if the user exists already, get user
	var _user user
	if sesh, ok := dbSessions[cookie.Value]; ok {
		sesh.lastActivity = time.Now()
		_user = dbUsers[sesh.username]
	}
	return _user
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}

	sesh, ok := dbSessions[cookie.Value]
	if ok {
		sesh.lastActivity = time.Now()
		dbSessions[cookie.Value] = sesh

	}
	_, ok = dbUsers[sesh.username]
	//refresh session
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)
	return ok
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN!") //for demo
	showSessions()               //for demo

	for key, val := range dbSessions {
		if time.Since(val.lastActivity) > (time.Second * 30) {
			delete(dbSessions, key)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") //for demo
	showSessions()             //for demo

}

//for demo
func showSessions() {
	fmt.Println("****************")
	for key, val := range dbSessions {
		fmt.Println(key, val.username)
	}
	fmt.Println("")
}

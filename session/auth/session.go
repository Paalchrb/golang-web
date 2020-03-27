package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(w, cookie)

	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")

	if err != nil {
		return false
	}

	un := dbSessions[cookie.Value]
	_, ok := dbUsers[un]
	return ok
}
	
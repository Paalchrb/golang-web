package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<h2><a href="/set">Set cookie</a></h2>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "my-cookie",
		Value: "This is a test cookie",
	})

	fmt.Fprintln(w, `<h1>Cookie set successfully!</h1>`)
	fmt.Fprintln(w, `<h2><a href="/read">Read cookie</a></h2>`)
}

func read(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintln(w, `<h1>Cookie read successfully</h1>`)
	fmt.Fprintln(w, `<h2>`, cookie, `</h2>`)
	fmt.Fprintln(w, `<h2><a href="/expire">Clear cookie</a></h2>`)
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1 // delete cookie
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
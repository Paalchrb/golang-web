package main

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, e := uuid.NewV4()
		if e != nil {
			log.Fatalln(e)
			return
		}

		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}

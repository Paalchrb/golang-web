package main

import (
	"fmt"
	"net/http"
)

func home (w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to homepage</h1>")
}

func dog (w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>I'm a dog. WOOOF WOOOF")
}

func me (w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>My name is Paal<h1>")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
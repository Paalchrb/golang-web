package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("*.gohtml"))
}

func home (w http.ResponseWriter, req *http.Request) {
	err := tmp.ExecuteTemplate(w, "home.gohtml", nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func dog (w http.ResponseWriter, req *http.Request) {
	err := tmp.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func me (w http.ResponseWriter, req *http.Request) {
	err := tmp.ExecuteTemplate(w, "me.gohtml", "Paal")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
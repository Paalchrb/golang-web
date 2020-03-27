package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmp, err := template.ParseFiles("templ.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

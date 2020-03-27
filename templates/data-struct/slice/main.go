package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	people := []string{"Peter", "Mary", "Jonah", "Gina"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", people)
	if err != nil {
		log.Fatalln(err)
	}
}

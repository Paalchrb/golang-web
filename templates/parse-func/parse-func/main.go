package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

// create a FuncMap to register functions.
// "uc" is what the function will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a function i declared
// "ft" slices a string returning the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type sage struct{
	Name string
	Motto string
}

func main() {
	buddha := sage{
		Name: "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name: "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name: "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed",
	}

	sages := []sage{buddha, gandhi, mlk}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
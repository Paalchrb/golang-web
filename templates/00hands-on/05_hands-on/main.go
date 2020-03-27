package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

type item struct{
	Name string
	Price float64
}

type meal struct{
	Meal string
	Items []item
}

type menu []meal

func main() {
	m := menu{
		meal{
			Meal: "Breakfast",
			Items: []item{
				item{"Bacon and eggs", 12.50},
				item{"Omelette", 10.25},
			},
		},
		meal{
			Meal: "Lunch",
			Items: []item{
				item{"Sandwich", 13.50},
				item{"Cheeseburger", 11.25},
			},
		},
		meal{
			Meal: "Dinner",
			Items: []item{
				item{"Beef Wellington", 22.70},
				item{"Pizza Margarita", 14.45},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
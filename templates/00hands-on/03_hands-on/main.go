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

type hotel struct{
	Name string
	Address string
	City string
	Zip string
}

type region struct{
	Region string
	Hotels []hotel
}

func main() {
	h := region{
		Region: "Southern",
		Hotels: []hotel{
			hotel{
				Name:    "Hotel California",
				Address: "42 Sunset Boulevard",
				City:    "Los Angeles",
				Zip:     "95612",
			},
			hotel{
				Name:    "H",
				Address: "4",
				City:    "L",
				Zip:     "95612",
			},
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
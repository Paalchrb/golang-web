package main

import "fmt"

type person struct{
	fn string
	ln string
	age int
}

type secretAgent struct{
	person
	ltk bool
}

func (sa secretAgent) saSpeak() {
	fmt.Println("My name is ", sa.ln, ", ", sa.fn, sa.ln, "and I've got licence to kill: ", sa.ltk)
}

func (p person) pSpeak() {
	fmt.Println(p.fn, p.ln ,"says hello! My age is ", p.age)
}

func main() {
	p1 := person{
		"Peter",
		"Pan",
		11,
	}

	sa := secretAgent{
		person{
			"James",
			"bond",
			39,
		},
		true,
	}

	fmt.Println("#### Person ####")
	fmt.Println(p1.fn)
	p1.pSpeak()
	
	fmt.Println("#### Secret Agent ####")
	fmt.Println(sa.ltk)
	sa.saSpeak()
	sa.person.pSpeak()

}
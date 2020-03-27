package main

import "fmt"

type person struct{
	fName string
	lName string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fName, "is walking")
}

func main() {
	p1 := person{
		"Peter",
		"Pan",
		[]string{"pasta", "pizza", "lasagna"},
	}

	fmt.Println(p1)
	fmt.Println(p1.fName)
	fmt.Println(p1.favFood)

	for i, v := range p1.favFood {
		fmt.Println("index:", i, "\tvalue:", v)
	}

	s := p1.walk()
	fmt.Println(s)
}
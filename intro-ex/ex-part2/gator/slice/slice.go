package main

import "fmt"

func main() {
	xi := []int{2,3,4,5,6,7,8,9}
	
	fmt.Printf("%v is type %T\n", xi, xi)

	for i := range xi {
		fmt.Println("Index is:", i)
	}
	
	for i, v := range xi {
		fmt.Println("Index is:", i, "and value is:", v)
	}
}
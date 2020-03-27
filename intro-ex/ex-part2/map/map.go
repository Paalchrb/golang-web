package main

import "fmt"

func main() {
	m := map[string]int{
		"Peter": 45,
		"Joe":   12,
		"Mary":  24,
	}

	fmt.Printf("value: %v\t type: %T\n", m, m)

	for k := range m {
		fmt.Println("key:", k)
	}

	for k, v := range m {
		fmt.Println("key:", k, "\tvalue:", v)
	}
}

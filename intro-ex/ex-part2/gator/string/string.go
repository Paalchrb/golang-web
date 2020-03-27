package main

import "fmt"

func main() {
	s := "I'm sorry Dave, I can't do that!"

	fmt.Println(s)
	fmt.Println([]byte(s))
	fmt.Println(string([]byte(s)))
	fmt.Println(s[:14])
	fmt.Println(s[10:23])
	fmt.Println(s[18:])

	for _, v := range []byte(s) {
		fmt.Println(string(v))
	}
}

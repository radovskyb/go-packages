package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new string s
	s := "hello, world. how are you?"

	// Use bytes.Title to Titleise each word in string s
	//
	// Title returns a copy of s with all Unicode
	// letters that begin words mapped to their title case.
	st := bytes.Title([]byte(s))

	// Print out the results of bytes.Title on s as a string
	// Prints: Hello, World. How Are You?
	fmt.Println(string(st))
}

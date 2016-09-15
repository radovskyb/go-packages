package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "Hello"
	b := "Hello"
	c := "World"

	// Compare if strings a and b match
	//
	// Compare returns an integer comparing two strings lexicographically.
	// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
	//
	// Compare is included only for symmetry with package bytes.
	// It is usually clearer and always faster to use the built-in
	// string comparison operators ==, <, >, and so on.
	if strings.Compare(a, b) == 0 {
		fmt.Println("a and b match")
	} else {
		fmt.Println("a and b don't match")
	}

	// Compare if strings a and c match
	if strings.Compare(a, c) == 0 {
		fmt.Println("a and c match")
	} else {
		fmt.Println("a and c don't match")
	}
}

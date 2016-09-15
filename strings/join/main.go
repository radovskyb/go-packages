package main

import (
	"fmt"
	"strings"
)

func main() {
	// Join two strings with a comma `,` and print the newly
	// created concatenated string
	//
	// Join concatenates the elements of a to create a single
	// string. The separator string sep is placed between
	// elements in the resulting string.
	fmt.Println(strings.Join([]string{"Hello", "World!"}, ", "))
}

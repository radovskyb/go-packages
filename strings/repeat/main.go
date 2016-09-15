package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a new string called `hello`
	hello := "Hello, World!\n"

	// Pass hello into the function strings.Repeat and pass as the count
	// parameter `2` which will print out the string `hello` twice
	//
	// Repeat returns a new string consisting of count copies
	// of the string s.
	fmt.Println(strings.Repeat(hello, 2))

	// Banana example from golang.org
	fmt.Println("ba" + strings.Repeat("na", 2))
}

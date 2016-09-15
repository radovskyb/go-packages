package main

import (
	"fmt"
	"strings"
)

func main() {
	// ToTitle returns a copy of the string s with all
	// Unicode letters mapped to their title case.
	fmt.Println(strings.ToTitle("hello, world!"))
}

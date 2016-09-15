package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a new string str
	str := "   Hello   "

	// Print out str before calling strings.Trim for comparison
	fmt.Println(str)

	// Trim all space from either side of the string str
	//
	// Trim returns a slice of the string s with all leading and
	// trailing Unicode code points contained in cutset removed.
	fmt.Println(strings.Trim(str, " "))
}

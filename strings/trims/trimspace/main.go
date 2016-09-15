package main

import (
	"fmt"
	"strings"
)

func main() {
	// Trim any leading or trailing white space characters from the
	// string below, which also include \n, \r and \t characters
	//
	// TrimSpace returns a slice of the string s, with all leading
	// and trailing white space removed, as defined by Unicode.
	fmt.Println(strings.TrimSpace("\n\n   Hello, World!   \n\n"))
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Check whether both the below strings can be backquoted
	//
	// CanBackquote reports whether the string s can be represented
	// unchanged as a single-line backquoted string without control
	// characters other than tab.
	fmt.Println(strconv.CanBackquote("I'm a string that can be backquoted"))
	fmt.Println(strconv.CanBackquote("`can't backquote this...`"))
}

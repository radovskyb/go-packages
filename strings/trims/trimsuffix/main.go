package main

import (
	"fmt"
	"strings"
)

func main() {
	// Trim the suffix `!` from the string below
	//
	// TrimSuffix returns s without the provided trailing suffix string.
	// If s doesn't end with suffix, s is returned unchanged.
	fmt.Println(strings.TrimSuffix("Hello, World!", "!"))
}

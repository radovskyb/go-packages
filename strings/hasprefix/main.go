package main

import (
	"fmt"
	"strings"
)

func main() {
	// HasPrefix tests whether the string s begins with prefix.
	fmt.Println(strings.HasPrefix("hello world", "hello")) // true
	fmt.Println(strings.HasPrefix("hello world", "ello"))  // false
	fmt.Println(strings.HasPrefix("hello world", ""))      // true
}

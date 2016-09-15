package main

import (
	"fmt"
	"strings"
)

func main() {
	// HasSuffix tests whether the string s ends with suffix.
	fmt.Println(strings.HasSuffix("hello world", "world"))  // true
	fmt.Println(strings.HasSuffix("hello world", "worldy")) // false
	fmt.Println(strings.HasSuffix("hello world", ""))       // true
}

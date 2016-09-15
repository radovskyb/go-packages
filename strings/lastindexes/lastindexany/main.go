package main

import (
	"fmt"
	"strings"
)

func main() {

	// LastIndexAny returns the index of the last instance
	// of any Unicode code point from chars in s, or -1 if no
	// Unicode code point from chars is present in s.
	fmt.Println(strings.LastIndexAny("Hello, World!", "abcd")) // 11
	fmt.Println(strings.LastIndexAny("Hello, World!", "xyz"))  // -1
}

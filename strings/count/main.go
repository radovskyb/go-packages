package main

import (
	"fmt"
	"strings"
)

func main() {
	// Count how many of the letter h are in the string `hip-hop`
	//
	// Count counts the number of non-overlapping instances
	// of sep in s. If sep is an empty string, Count returns
	// 1 + the number of Unicode code points in s.
	fmt.Println(strings.Count("hip-hop", "h")) // 2

	// Count how many of the letter z are in the string `hip-hop`
	fmt.Println(strings.Count("hip-hop", "z")) // 0

	// Before & after each rune
	fmt.Println(strings.Count("five", "")) // 5
}

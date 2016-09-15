package main

import (
	"fmt"
	"strings"
)

func main() {
	// Split the string "Hello World" by spaces
	//
	// Split slices s into all substrings separated by sep and
	// returns a slice of the substrings between those separators.
	// If sep is empty, Split splits after each UTF-8 sequence.
	// It is equivalent to SplitN with a count of -1.
	split := strings.Split("Hello World", " ")

	// Print out the split pieces
	for _, str := range split {
		fmt.Println(str)
	}

	// This time strings.Split will split between each UTF-8
	// character since sep is a blank string
	split = strings.Split("Hello World", "")
	fmt.Println(split)
}

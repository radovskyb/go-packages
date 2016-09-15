package main

import (
	"fmt"
	"strings"
)

func main() {
	// Split the string below into 2 strings by the separator `,`
	//
	// SplitN slices s into substrings separated by sep and
	// returns a slice of the substrings between those separators.
	// If sep is empty, SplitN splits after each UTF-8 sequence.
	// The count determines the number of substrings to return:
	//   n > 0: at most n substrings; the last substring will be
	//			the unsplit remainder.
	//   n == 0: the result is nil (zero substrings)
	//   n < 0: all substrings
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))

	// Split and return no substrings since n int is 0
	fmt.Println(strings.SplitN("a,b,c", ",", 0))

}

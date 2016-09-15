package main

import (
	"fmt"
	"strings"
)

func main() {
	// TrimLeft returns a slice of the string s with all leading
	// Unicode code points contained in cutset removed.
	fmt.Println(strings.TrimLeft("Hello, World!", "Hello, ")) // World!

	// TrimRight returns a slice of the string s, with all trailing
	// Unicode code points contained in cutset removed.
	fmt.Println(strings.TrimRight("Hello, World!", ", World!")) // He
}

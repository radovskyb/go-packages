package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	// Create a new strings reader for the string below
	sr := strings.NewReader("Hello, World!")

	// Read 1 byte from sr into `b` and it's size into `size`
	r, size, err := sr.ReadRune()
	if err != nil {
		log.Fatalln(err)
	}

	// Print out `r` as a string and `r`'s size from `size`
	fmt.Printf("%v: %s", size, string(r)) // H
}

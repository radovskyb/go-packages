package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	// Create a new strings reader for the string below
	sr := strings.NewReader("Hello, World!")

	// Read 1 byte from sr into `r`
	r, _, err := sr.ReadRune()
	if err != nil {
		log.Fatalln(err)
	}

	// Print out `r` as a string
	fmt.Println(string(r)) // H

	// Print out the amount of unread bytes in `sr`
	fmt.Println(sr.Len()) // 12

	// Now unread the previously read rune `H`
	sr.UnreadRune()

	// Once again print out the amount of unread byte in `sr`
	fmt.Println(sr.Len()) // 13
}

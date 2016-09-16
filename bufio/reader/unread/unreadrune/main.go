package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World!"

	// Create a new buffered reader which reads from
	// a string reader that reads from string s
	br := bufio.NewReader(strings.NewReader(s))

	// Read one rune into b from br
	b, _, err := br.ReadRune()
	if err != nil {
		log.Fatalln(err)
	}

	// Print that rune
	fmt.Println(string(b)) // H

	// Now unread the previous rune read in from br.
	// Instead of next time reading in rune `e`, br will
	// once again return rune `H`
	//
	// UnreadRune unreads the last rune. If the most recent read
	// operation on the buffer was not a ReadRune, UnreadRune returns
	// an error. (In this regard it is stricter than UnreadByte, which
	// will unread the last byte from any read operation.)
	if err := br.UnreadRune(); err != nil {
		log.Fatalln(err)
	}

	// Read a rune from br again
	b, _, err = br.ReadRune() // `H` read instead of `e`
	if err != nil {
		log.Fatalln(err)
	}

	// Print the rune again, however
	// this time it will be back at rune `H`
	// instead of rune `e`
	fmt.Println(string(b)) // H
}

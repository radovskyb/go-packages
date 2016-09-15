package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	// Create a new strings reader for the string below
	sr := strings.NewReader("Hello, World!")

	// Read 1 byte from sr into `b`
	b, err := sr.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}

	// Print out `b` as a string
	fmt.Println(string(b)) // H
}

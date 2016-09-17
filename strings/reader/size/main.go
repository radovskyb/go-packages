package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	// Create a new strings reader for the string below
	sr := strings.NewReader("Hello, World!")

	// Create a byte slice `buf` with the length of
	// the unread portion of the strings reader `sr`
	buf := make([]byte, sr.Len())

	// Read from sr into buf
	_, err := sr.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}

	// Print out buf as a string
	fmt.Println(string(buf)) // Hello, World!

	// Print out the amount of unread bytes in strings reader `sr`
	fmt.Println(sr.Len()) // 0

	// Print out the original length of the underlying string in `sr`
	//
	// Size returns the original length of the underlying string.
	// Size is the number of bytes available for reading via ReadAt.
	// The returned value is always the same and is not affected by calls
	// to any other method.
	fmt.Println(sr.Size()) // 13
}

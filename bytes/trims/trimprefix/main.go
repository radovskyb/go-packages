package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	// Declare variable b which is a byte array of
	// the string `Hello, World!`
	var b = []byte("Hello, World!")

	b = bytes.TrimPrefix(b, []byte("Hello, "))

	// No change will be made here as `Gopher` is
	// not inside the byte slice b
	b = bytes.TrimPrefix(b, []byte("Gopher!"))

	// Write b to os.Stdout
	_, err := os.Stdout.Write(b)
	if err != nil {
		log.Fatalln(err)
	}

	// Write a blank line
	fmt.Println()

	// Now append the byte slice `!Hello, ` to the beginning of b,
	// but first use bytes.TrimPrefix to trim any `!` prefix from `!Hello, `
	b = append(bytes.TrimPrefix([]byte("!Hello, "), []byte("!")), b...)

	// Write b to os.Stdout again
	_, err = os.Stdout.Write(b)
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// Create a new string s
	s := "Hello, World"

	// Create a new buffered writer to write to os.Stdout
	w := bufio.NewWriter(os.Stdout)

	// Use w.WriteRune to write a single rune into w's buffer
	//
	// WriteRune writes a single Unicode code point, returning
	// the number of bytes written and any error.
	_, err := w.WriteRune([]rune(s)[0])
	if err != nil {
		log.Fatalln(err)
	}

	// Flush w to actually write w's contents to os.Stdout
	if err := w.Flush(); err != nil {
		log.Fatalln(err)
	}
}

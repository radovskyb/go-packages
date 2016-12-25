package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Create some sample artificial input to be used to write to a writer
	const input = "This is some sample input to be written to a writer"

	// Create a new buffered writer to os.Stdout using bufio.NewWriter
	w := bufio.NewWriter(os.Stdout)

	// Use Fprintln to write the input const string to our
	// buffered writer w which is set to write to os.Stdout
	fmt.Fprintln(w, input)

	// We now flush the buffered writer w, using w.Flush()
	// Flush writes any buffered data to the underlying io.Writer
	// by actually calling the writers Write() method
	if err := w.Flush(); err != nil {
		log.Fatalln(err)
	}
}

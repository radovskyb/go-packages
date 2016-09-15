package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create some sample artifical input to be used to write to a writer
	const input = "This is some sample input to be written to a writer"

	// Create a new buffered writer to os.Stdout using bufio.NewWriter
	w := bufio.NewWriter(os.Stdout)

	// Read the contents from input into buffered writer w
	//
	// ReadFrom implements io.ReaderFrom.
	w.ReadFrom(strings.NewReader(input))

	// Print from buffered writer w to os.Stdout
	fmt.Fprintln(w)

	// We now flush the buffered writer w, using w.Flush()
	// Flush writes any buffered data to the underlying io.Writer
	// by actually calling the writers Write() method
	w.Flush()
}

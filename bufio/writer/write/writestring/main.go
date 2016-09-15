package main

import (
	"bufio"
	"os"
)

func main() {
	// Create a new string s
	s := "Hello, World"

	// Create a new buffered writer to write to os.Stdout
	w := bufio.NewWriter(os.Stdout)

	// Use w.WriteString to write string s into w's buffer
	//
	// WriteString writes a string. It returns the number of
	// bytes written. If the count is less than len(s), it also
	// returns an error explaining why the write is short.
	w.WriteString(s)

	// Flush w to actually write w's contents to os.Stdout
	w.Flush()
}

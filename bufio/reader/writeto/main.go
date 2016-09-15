package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello"

	// Create a new buffered reader br that reads from
	// a new strings reader that reads from string s
	br := bufio.NewReader(strings.NewReader(s))

	// Write the buffered reader br's contents to os.Stdout
	//
	// WriteTo implements io.WriterTo.
	br.WriteTo(os.Stdout)
}

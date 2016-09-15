package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World!" // 13 bytes long

	// Create a new buffered reader br that will
	// read from a new strings reader for string s
	br := bufio.NewReader(strings.NewReader(s))

	// Read the first 6 bytes up to the first comma
	// using ReadString(delim byte)
	br.ReadString(byte(',')) // 7 bytes left

	// br now contains ` World!`
	// Discard the next 6 bytes from br
	// Since ` World!` is 7 bytes long, we
	// will only be left with 1 byte of `!`
	br.Discard(6) // br now only contains `!`

	// Print out the remaining string
	// from the buffered reader br
	io.Copy(os.Stdout, br)
}

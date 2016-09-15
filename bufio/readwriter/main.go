package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World"

	// Create a new buffered reader which takes a new
	// strings reader for s as it's reader parameter
	r := bufio.NewReader(strings.NewReader(s))

	// Create a new buffered writer for os.Stdout
	w := bufio.NewWriter(os.Stdout)

	// Create a new buffered ReadWriter rw, that takes
	// r reader and w writer as its parameters.
	rw := bufio.NewReadWriter(r, w)

	// Use io.Copy to copy from rw to rw, which will use
	// rw's buffered writer and rw's buffered reader
	// as its parameters respectively
	io.Copy(rw, rw)

	// Now to make sure that w actually writes to os.Stdout,
	// we need to call Flush() on rw which will push anything
	// inside of rw's writer to os.Stdout
	rw.Flush()
}

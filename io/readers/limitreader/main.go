package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Create a new limited reader that reads from stdin.
	// This is the same as actually calling:
	// r := io.LimitedReader{os.Stdin, 10}
	//
	// LimitReader returns a Reader that reads from r but stops with
	// EOF after n bytes. The underlying implementation is a *LimitedReader.
	r := io.LimitReader(os.Stdin, 10)

	// Create a new buffer to read into
	buf := make([]byte, 10)

	// Read from r into buf
	r.Read(buf)

	// Now print the contents of buf
	fmt.Println(string(buf))
}

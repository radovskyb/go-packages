package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new string s
	s := "Hello, World!"

	// Create a new bytes buffer buf with s for it's initial content
	buf := bytes.NewBuffer([]byte(s))

	// Create a new byte slice b of len 13 to
	// read from our buffer into
	b := make([]byte, 13)

	// Read from buffer buf into byte slice b
	//
	// Read reads the next len(p) bytes from the buffer or until
	// the buffer is drained. The return value n is the number of
	// bytes read. If the buffer has no data to return, err is
	// io.EOF (unless len(p) is zero); otherwise it is nil.
	buf.Read(b)

	// Print out the results of buf.Read into byte slice b
	fmt.Println(string(b))
}

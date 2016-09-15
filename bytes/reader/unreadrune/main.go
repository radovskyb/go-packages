package main

import (
	"bytes"
	"os"
)

func main() {
	// Create a new bytes reader with some contents
	r := bytes.NewReader([]byte{'a', 'b', 'c'})

	// Read one rune from r
	r.ReadRune() // r's contents are now `bc`

	// Now unread the byte that was read before
	r.UnreadRune() // r's contents are back to `abc`

	// Print r's contents
	r.WriteTo(os.Stdout) // abc
}

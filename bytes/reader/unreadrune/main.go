package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	// Create a new bytes reader with some contents
	r := bytes.NewReader([]byte{'a', 'b', 'c'})

	// Read one rune from r
	_, _, err := r.ReadRune() // r's contents are now `bc`
	if err != nil {
		log.Fatalln(err)
	}

	// Now unread the byte that was read before
	if err := r.UnreadRune(); err != nil { // r's contents are back to `abc`
		log.Fatalln(err)
	}

	// Print r's contents
	_, err = r.WriteTo(os.Stdout) // abc
	if err != nil {
		log.Fatalln(err)
	}
}

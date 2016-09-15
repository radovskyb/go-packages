package main

import (
	"bytes"
	"os"
)

func main() {
	// Create a new bytes reader r with some contents
	r := bytes.NewReader([]byte{'a', 'b', 'c'})

	// Write r's contents to os.Stdout
	r.WriteTo(os.Stdout)
}

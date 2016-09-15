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

	// Use w.WriteByte to write a single byte into w's buffer
	//
	// WriteByte writes a single byte.
	w.WriteByte([]byte(s)[0])

	// Flush w to actually write w's contents to os.Stdout
	w.Flush()
}

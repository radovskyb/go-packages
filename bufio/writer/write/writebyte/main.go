package main

import (
	"bufio"
	"log"
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
	if err := w.WriteByte([]byte(s)[0]); err != nil {
		log.Fatalln(err)
	}

	// Flush w to actually write w's contents to os.Stdout
	if err := w.Flush(); err != nil {
		log.Fatalln(err)
	}
}

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

	// Write the contents of string s into w's writer's buffer
	//
	// Write writes the contents of p into the buffer. It returns
	// the number of bytes written. If nn < len(p), it also returns
	// an error explaining why the write is short.
	_, err := w.Write([]byte(s))
	if err != nil {
		log.Fatalln(err)
	}

	// Flush w to actually write w's contents to os.Stdout
	if err := w.Flush(); err != nil {
		log.Fatalln(err)
	}
}

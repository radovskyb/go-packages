package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create a new pipe reader and writer
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatalln(err)
	}

	// Write to the pipe writer.
	_, err = w.Write([]byte("Hello, World!"))
	if err != nil {
		log.Fatalln(err)
	}

	// Crete a new string channel that will receive data from the pipe reader
	// from inside of the go routine below
	data := make(chan string)

	// Create a new go routine
	go func() {
		// Create a new bytes buffer that will store data from the pipe reader
		var buf bytes.Buffer

		// Copy the data from the pipe reader `r` into the bytes buffer `buf`
		_, err := io.Copy(&buf, r)
		if err != nil {
			log.Fatalln(err)
		}

		// Now pass the buffer buf's contents as a string to the data channel
		data <- buf.String()
	}()

	// Close the pipe writer to allow the data to pass through so we do not
	// continually block
	if err := w.Close(); err != nil {
		log.Fatalln(err)
	}

	// Print out the data received on the channel data
	fmt.Println(<-data)
}

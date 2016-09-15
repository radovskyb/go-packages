package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// Create a new pipe reader and writer
	r, w, _ := os.Pipe()

	// Writer to the pipe writer through the function Write, which in reality
	// could process data in any way. (Here we just pass it data to write)
	Write(w, "Hello, World!")

	// Crete a new string channel that will receive data from the pipe reader
	// from inside of the go routine below
	data := make(chan string)

	// Create a new go routine
	go func() {
		// Create a new bytes buffer that will store data from the pipe reader
		var buf bytes.Buffer

		// Copy the data from the pipe reader `r` into the bytes buffer `buf`
		io.Copy(&buf, r)

		// Now pass the buffer buf's contents as a string to the data channel
		data <- buf.String()
	}()

	// Close the pipe writer to allow the data to pass through so we do not
	// continually block
	w.Close()

	// Print out the data received on the channel data
	fmt.Println(<-data)
}

func Write(w *os.File, data string) {
	w.Write([]byte(data))
}

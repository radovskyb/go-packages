package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a new pipe reader and pipe writer, r and w, respectively.
	//
	// This example creates a pipe reader and writer, then passes that writer
	// to a function, which writes some data it it, then once the function
	// has finished executing, outside of the function, the pipe reader reads
	// whatever it has received from the writer, into a data byte slice `data`
	// and then print what is now stored inside of data from the pipe reader.
	//
	// Pipe returns a connected pair of Files; reads from r return bytes
	// written to w. It returns the files and an error, if any.
	r, w, _ := os.Pipe()

	// Pass the pipe writer to a function that writes data to it,
	// in this case, getting it's data from its string parameter data
	n, err := WriteSomeDataToWriter(w, "Hello, World!")

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Create a new byte slice `data` that is the size of the data written
	// to w, returned as n from the function `WriteSomeDataToWriter`
	data := make([]byte, n)

	// Read the data from the pipe reader into byte slice `data`
	r.Read(data)

	// Print out the byte slice `data` as a string
	fmt.Println(string(data))
}

func WriteSomeDataToWriter(w *os.File, data string) (n int, err error) {
	n, err = w.Write([]byte(data))
	return
}

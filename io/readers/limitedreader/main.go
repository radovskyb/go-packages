package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create a new limited reader that has 5 bytes
	// for its max bytes remaining field
	//
	// A LimitedReader reads from R but limits the amount of data
	// returned to just N bytes. Each call to Read updates N to
	// reflect the new amount remaining.
	lr := io.LimitedReader{os.Stdin, 5}

	// Create a new byte slice buffer buf of size 5 bytes
	buf := make([]byte, 5)

	// Iterate 10 times
	for i := 0; i < 10; i++ {
		// Read using the limited reader into buf
		//
		// A LimitedReader reads from R but limits the amount of data
		// returned to just N bytes. Each call to Read updates N to
		// reflect the new amount remaining.
		n, err := lr.Read(buf[0:])

		// Log any errors that may have occurred
		if err != nil {
			log.Fatalln(err)
		}

		// Reset lr.N to 5 so we can read again
		lr.N = 5

		// Print out how many bytes were read using the limited reader
		fmt.Printf("Read %d bytes from the limited reader\nBytes read: %s\n", n, buf)
	}
}

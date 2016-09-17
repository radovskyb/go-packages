package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Create a new TeeReader that reads from a string reader that reads
	// from the string `Hello, World!`
	//
	// TeeReader returns a Reader that writes to w what it reads from r.
	// All reads from r performed through it are matched with corresponding
	// writes to w. There is no internal buffering - the write must
	// complete before the read completes. Any error encountered while
	// writing is reported as a read error.
	tr := io.TeeReader(strings.NewReader("Hello, World!\n"), os.Stdout)

	// Create a new byte slice variable buf to hold reads from
	// TeeReader tr
	buf := make([]byte, 14)

	// Read from tr into buf
	// This will now print what it reads to os.Stdout as
	// os.Stdout was used as TeaReaders writer parameter
	_, err := tr.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
}

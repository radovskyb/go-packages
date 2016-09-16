package main

import (
	"compress/lzw"
	"fmt"
	"log"
	"os"
)

// Create a string str to compress into lzw format
const str = "Hello, World!\n"

func main() {
	// Create a new file to store the compressed version of str
	fout, err := os.Create("hello.txt.lzw")
	if err != nil {
		log.Fatalln(err)
	}
	defer fout.Close()

	// Create a new lzw writer that writes to fout
	//
	// NewWriter creates a new io.WriteCloser.
	// Writes to the returned io.WriteCloser are compressed and written to w.
	// It is the caller's responsibility to call Close on the WriteCloser when
	// finished writing.
	// The number of bits to use for literal codes, litWidth, must be in the
	// range [2,8] and is typically 8. Input bytes must be less
	// than 1<<litWidth.
	lzwWriter := lzw.NewWriter(fout, lzw.LSB, 8)

	// Write str to lzwWriter
	n, err := lzwWriter.Write([]byte(str))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Wrote", n, "bytes to fout")

	// Close lzwWriter
	if err := lzwWriter.Close(); err != nil {
		log.Fatalln(err)
	}
}

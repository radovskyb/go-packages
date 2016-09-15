package main

import (
	"compress/lzw"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open hello.txt.lzw for reading
	fin, err := os.Open("hello.txt.lzw")
	if err != nil {
		log.Fatalln(err)
	}
	defer fin.Close()

	// Create a new lzw reader that reads data from fin in
	// a decompressed format
	//
	// NewReader creates a new io.ReadCloser.
	// Reads from the returned io.ReadCloser read and decompress data from r.
	// If r does not also implement io.ByteReader,
	// the decompressor may read more data than necessary from r.
	// It is the caller's responsibility to call Close on the ReadCloser when
	// finished reading.
	// The number of bits to use for literal codes, litWidth, must be in the
	// range [2,8] and is typically 8. It must equal the litWidth
	// used during compression.
	lzwReader := lzw.NewReader(fin, lzw.LSB, 8)
	defer lzwReader.Close()

	// Read the data from lzwReader and write it to os.Stdout
	n, err := io.Copy(os.Stdout, lzwReader)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Read", n, "bytes from lzwReader")
}

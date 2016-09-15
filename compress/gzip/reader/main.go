package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open file.gz for reading the compressed gzip data from
	fin, err := os.Open("file.gz")

	// Create a new gzip reader to read the compressed data from fin
	//
	// NewReader creates a new Reader reading the given reader.
	// If r does not also implement io.ByteReader,
	// the decompressor may read more data than necessary from r.
	//
	// It is the caller's responsibility to call Close on the Reader when done.
	//
	// The Reader.Header fields will be valid in the Reader returned.
	gzipReader, err := gzip.NewReader(fin)
	if err != nil {
		log.Fatalln(err)
	}
	defer gzipReader.Close()

	// Read the data in uncompressed gzip format to os.Stdout
	// from gzipReader, decompressing from &buf
	n, err := io.Copy(os.Stdout, gzipReader)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nRead", n, "bytes from gzipReader")
}

package main

import (
	"compress/flate"
	"io"
	"log"
	"os"
)

var filename = "hello.txt.compressed"

func main() {
	// Open up the flate compressed file hello.txt.compressed
	fin, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer fin.Close()

	// Create a new flate reader to read compressed data from fin
	//
	// NewReader returns a new ReadCloser that can be used
	// to read the uncompressed version of r.
	// If r does not also implement io.ByteReader,
	// the decompressor may read more data than necessary from r.
	// It is the caller's responsibility to call Close on the ReadCloser
	// when finished reading.
	//
	// The ReadCloser returned by NewReader also implements Resetter.
	flateReader := flate.NewReader(fin)
	defer flateReader.Close()

	// Copy the compressed data from fin, to os.Stdout in
	// non-compressed format.
	_, err = io.Copy(os.Stdout, flateReader)
	if err != nil {
		log.Fatalln(err)
	}
}

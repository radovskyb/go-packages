package main

import (
	"compress/gzip"
	"fmt"
	"log"
	"os"
)

// Create a string to be compressed into gzip format
const str = `I'm a string. Hello, World!`

func main() {
	// Create a new file to compress str into gzip format into
	fout, err := os.Create("file.gz")
	if err != nil {
		log.Fatalln(err)
	}
	defer fout.Close()

	// NewWriter returns a new Writer.
	// Writes to the returned writer are compressed and written to w.
	//
	// It is the caller's responsibility to call Close on the
	// WriteCloser when done.
	//
	// Writes may be buffered and not flushed until Close.
	//
	// Callers that wish to set the fields in Writer.Header must do so before
	// the first call to Write, Flush, or Close.
	gzipWriter := gzip.NewWriter(fout)

	// Write the data from a strings reader reading
	// from str into gzipWriter
	// n, err := io.Copy(gzipWriter, strings.NewReader(str))
	n, err := gzipWriter.Write([]byte(str))
	if err != nil {
		log.Fatalln(err)
	}

	// Close closes the Writer, flushing any unwritten data to the underlying
	// io.Writer, but does not close the underlying io.Writer.
	if err := gzipWriter.Close(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Compressed", n, "bytes into fout")
}

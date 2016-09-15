package main

import (
	"compress/bzip2"
	"fmt"
	"io"
	"log"
	"os"
)

const filename = "hello.txt.bz2"

func main() {
	// Open a compressed bzip2 file
	fin, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer fin.Close()

	// Create a new file to store the a decompressed file from fin
	fout, err := os.Create("hello.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fout.Close()

	fmt.Println("Created fout file.")

	// Create a new bzip2 reader that reads from fin
	//
	// NewReader returns an io.Reader which decompresses bzip2 data from r.
	// If r does not also implement io.ByteReader, the decompressor
	// may read more data than necessary from r.
	bzip2reader := bzip2.NewReader(fin)

	// Copy the contents from bzip2reader into fout
	n, err := io.Copy(fout, bzip2reader)
	if err != nil {
		log.Fatalln("Erorr copying data from bzip2reader to fout file")
	}

	fmt.Println("Copied", n, "bytes from bzip2reader to fout file.")
}

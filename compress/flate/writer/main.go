package main

import (
	"compress/flate"
	"fmt"
	"io"
	"log"
	"os"
)

var filename = "hello.txt"

func main() {
	// Open the file hello.txt for reading
	fin, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer fin.Close()

	// Create a new file to store the compressed
	// version of hello.txt read in from fin
	fout, err := os.Create(filename + ".compressed")
	if err != nil {
		log.Fatalln(err)
	}
	defer fout.Close()

	// Create a new flate writer that writes to fout with the best level
	// of flate compression `flate.BestCompresion` set
	//
	// NewWriter returns a new Writer compressing data at the given level.
	// Following zlib, levels range from 1 (BestSpeed) to 9 (BestCompression);
	// higher levels typically run slower but compress more. Level 0
	// (NoCompression) does not attempt any compression; it only adds the
	// necessary DEFLATE framing. Level -1 (DefaultCompression) uses
	// the default compression level.
	//
	// If level is in the range [-1, 9] then the error returned will be nil.
	// Otherwise the error returned will be non-nil.
	flateWriter, err := flate.NewWriter(fout, flate.BestCompression)
	if err != nil {
		log.Fatalln(err)
	}
	defer flateWriter.Close()

	// Copy the data from fin into flateWriter
	n, err := io.Copy(flateWriter, fin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Compressed", n, "bytes")

	// Flush the data written to flateWriter
	//
	// Flush flushes any pending compressed data to the underlying writer.
	// It is useful mainly in compressed network protocols, to ensure that
	// a remote reader has enough data to reconstruct a packet.
	// Flush does not return until the data has been written.
	// If the underlying writer returns an error, Flush returns that error.
	//
	// In the terminology of the zlib library, Flush is equivalent
	// to Z_SYNC_FLUSH.
	err = flateWriter.Flush()
	if err != nil {
		log.Fatalln(err)
	}
}

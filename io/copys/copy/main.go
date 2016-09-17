package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// Copy from stdin to stdout (a simple echo program)
	//
	// Copy copies from src to dst until either EOF is reached on src
	// or an error occurs. It returns the number of bytes copied and
	// the first error encountered while copying, if any.
	//
	// A successful Copy returns err == nil, not err == EOF. Because
	// Copy is defined to read from src until EOF, it does not
	// treat an EOF from Read as an error to be reported.
	//
	// If src implements the WriterTo interface, the copy is implemented
	// by calling src.WriteTo(dst). Otherwise, if dst implements the
	// ReaderFrom interface, the copy is implemented by
	// calling dst.ReadFrom(src).
	_, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	// Create a new bytes buffer buf with some contents
	buf := bytes.NewBuffer([]byte{'a', 'b', 'c'})

	// Append a string to buf
	_, err := buf.WriteString("defghijk")
	if err != nil {
		log.Fatalln(err)
	}

	// Write buf to os.Stdout
	//
	// WriteTo writes data to w until the buffer is drained or an error
	// occurs. The return value n is the number of bytes written; it always
	// fits into an int, but it is int64 to match the io.WriterTo
	// interface. Any error encountered during the write is also returned.
	_, err = buf.WriteTo(os.Stdout) // abcdefghijk
	if err != nil {
		log.Fatalln(err)
	}
}

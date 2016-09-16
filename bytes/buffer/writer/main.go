package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a bytes.Buffer b
	var b bytes.Buffer // A Buffer needs no initialization

	// Write the string `Hello!` as a byte array
	// to the byte buffer b
	_, err := b.Write([]byte("Hello "))
	if err != nil {
		log.Fatalln(err)
	}

	// Now write the string "World!" into the byte buffer b,
	// using Fprintln, which accepts an io.Writer and a string
	// and writes to the given writer, in this case b
	fmt.Fprintln(&b, "World!")

	// Now we write the contents of b to os.Stdout using b
	// which is itself a writer so contains the method
	// WriteTo, which takes an io.Writer to write to
	_, err = b.WriteTo(os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}

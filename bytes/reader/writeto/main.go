package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a new bytes reader r with some contents
	r := bytes.NewReader([]byte{'a', 'b', 'c'})

	// Write r's contents to os.Stdout
	n, err := r.WriteTo(os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\nWrote %d bytes to os.Stdout\n", n)
}

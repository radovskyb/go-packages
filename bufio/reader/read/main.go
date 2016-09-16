package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World!"

	// Create a new buffered reader br which reads from
	// a new string reader that reads from the string s
	br := bufio.NewReader(strings.NewReader(s))

	// Create a new []byte read of the size len(s),
	// that will store the results of br.Read()
	read := make([]byte, len(s))

	// Read from buffered reader br into []byte read
	n, err := br.Read(read) // read is now `Hello, World!`
	if err != nil {
		log.Fatalln(err)
	}

	// Print []byte read as a string which will print
	// what was read using the buffered reader br
	fmt.Printf("Read %d bytes: %s\n", n, read)
}

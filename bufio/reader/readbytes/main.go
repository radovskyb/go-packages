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

	// Read any amount of bytes up until and including
	// the delimiter byte `,`
	//
	// ReadBytes reads until the first occurrence of delim in the
	// input, returning a slice containing the data up to and including
	// the delimiter. If ReadBytes encounters an error before finding a
	// delimiter, it returns the data read before the error and the error
	// itself (often io.EOF). ReadBytes returns err != nil if and only if
	// the returned data does not end in delim.
	//
	// For simple uses, a Scanner may be more convenient.
	b, err := br.ReadBytes(byte(','))
	if err != nil {
		log.Fatalln(err)
	}

	// Print the bytes read from br in string form
	fmt.Printf("Read %d bytes: %s\n", len(b), b) // Hello,
}

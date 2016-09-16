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

	// Read a and return a []byte pointing at the bytes in
	// the buffer up until the delimiter `,` (comma)
	//
	// ReadSlice reads until the first occurrence of delim in the input,
	// returning a slice pointing at the bytes in the buffer. The bytes
	// stop being valid at the next read. If ReadSlice encounters an error
	// before finding a delimiter, it returns all the data in the buffer and
	// the error itself (often io.EOF). ReadSlice fails with error ErrBufferFull
	// if the buffer fills without a delim. Because the data returned from
	// ReadSlice will be overwritten by the next I/O operation, most clients should
	// use ReadBytes or ReadString instead. ReadSlice returns err != nil if and
	// only if line does not end in delim.
	b, err := br.ReadSlice(byte(','))
	if err != nil {
		log.Fatalln(err)
	}

	// Print the byte slice read from br in string form
	fmt.Println("Read", len(b), "bytes:", string(b)) // Hello,
}

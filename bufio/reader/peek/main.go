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

	// Create a new buffered reader br that reads from
	// a new string reader that reads from string s
	br := bufio.NewReader(strings.NewReader(s))

	// Use br.Peek(5) to read the first 5 bytes from br
	// which will return a byte slice and the number
	// of elements peeked, in this case returning, `Hello`
	// and 5
	//
	// Peek returns the next n bytes without advancing the
	// reader. The bytes stop being valid at the next read call.
	// If Peek returns fewer than n bytes, it also returns an error
	// explaining why the read is short. The error is ErrBufferFull
	// if n is larger than b's buffer size.
	peeked, err := br.Peek(5)
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the byte slice returned from using
	// br.Peek(5), as a string
	fmt.Println(string(peeked))
}

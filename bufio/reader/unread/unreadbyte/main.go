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

	// Create a new buffered reader which reads from
	// a string reader that reads from string s
	br := bufio.NewReader(strings.NewReader(s))

	// Read one byte into b from br
	b, err := br.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}

	// Print that byte
	fmt.Println(string(b)) // H

	// Now unread the previous byte read in from br.
	// Instead of next time reading in byte `e`, br will
	// once again return byte `H`
	//
	// UnreadByte unreads the last byte. Only the most
	// recently read byte can be unread.
	if err := br.UnreadByte(); err != nil {
		log.Fatalln(err)
	}

	// Read a byte from br again
	b, err = br.ReadByte() // `H` read instead of `e`
	if err != nil {
		log.Fatalln(err)
	}

	// Print the byte again, however
	// this time it will be back at byte `H`
	// instead of byte `e`
	fmt.Println(string(b)) // H
}

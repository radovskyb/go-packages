package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World!"

	// Create a new buffered reader which reads from
	// a string reader that reads from string s
	br := bufio.NewReader(strings.NewReader(s))

	// Read one byte into b from br
	b, _ := br.ReadByte()

	// Print that byte
	fmt.Println(string(b)) // H

	// Now unread the previous byte read in from br.
	// Instead of next time reading in byte `e`, br will
	// once again return byte `H`
	//
	// UnreadByte unreads the last byte. Only the most
	// recently read byte can be unread.
	br.UnreadByte()

	// Read a byte from br again
	b, _ = br.ReadByte() // `H` read instead of `e`

	// Print the byte again, however
	// this time it will be back at byte `H`
	// instead of byte `e`
	fmt.Println(string(b)) // H
}

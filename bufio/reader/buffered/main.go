package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World!" // Length 13 bytes

	// Create a new buffered reader br, which will
	// read from a new strings reader for string s
	br := bufio.NewReader(strings.NewReader(s))

	// Read from br up until a comma in the string s
	br.ReadString(byte(',')) // Comma at byte 6

	// n = 13 bytes - 6 bytes
	n := br.Buffered() // n = 7

	// Print out how many bytes can
	// still be read from br
	fmt.Println(n) // 7
}

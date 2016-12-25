package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
)

func main() {
	// Create a new []byte and store it in variable b
	b := []byte("Hello, World!")

	// Create a new bytes reader for b called br
	br := bytes.NewReader(b)

	// Create a new buffered scanner using bufio.NewScanner
	// and pass it our bytes reader br
	scanner := bufio.NewScanner(br)

	// Set our scanner split function as bufio.ScanRunes
	//
	// ScanRunes is a split function for a Scanner that returns
	// each UTF-8-encoded rune as a token. The sequence of runes
	// returned is equivalent to that from a range loop over the
	// input as a string, which means that erroneous UTF-8 encodings
	// translate to U+FFFD = "\xef\xbf\xbd". Because of the Scan
	// interface, this makes it impossible for the client to distinguish
	// correctly encoded replacement runes from encoding errors.
	scanner.Split(bufio.ScanRunes)

	// Scan using our scanner until our scanning is complete
	for scanner.Scan() {
		// Print each token returned from scanner.Scan()
		// which is stored in scanner.Text()
		fmt.Println(scanner.Text())
	}

	// Check if any scanner errors occurred and if so log them
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

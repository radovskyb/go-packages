package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	// Create a new []byte and store it in variable b
	b := []byte("Hello, World!\nHow are we today?")

	// Create a new bytes reader for b called br
	br := bytes.NewReader(b)

	// Create a new buffered scanner using bufio.NewScanner
	// and pass it our bytes reader br
	scanner := bufio.NewScanner(br)

	// Set our scanner split function as bufio.ScanBytes
	//
	// ScanBytes is a split function for a Scanner that
	// returns each byte as a token.
	scanner.Split(bufio.ScanBytes)

	// Scan using our scanner until our scanning is complete
	for scanner.Scan() {
		// Print each token returned from scanner.Scan()
		// which is stored in scanner.Text()
		fmt.Println(scanner.Text())
	}
}

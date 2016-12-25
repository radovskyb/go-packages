package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	// Create a new string
	s := "Hello, World!"

	// Create a new strings reader for s
	sr := strings.NewReader(s)

	// Create a new buffered scanner for
	// string reader sr
	scanner := bufio.NewScanner(sr)

	// Set the scanner split function as bufio.ScanWords
	//
	// ScanWords is a split function for a Scanner that returns
	// each space-separated word of text, with surrounding spaces
	// deleted. It will never return an empty string. The definition
	// of space is set by unicode.IsSpace.
	scanner.Split(bufio.ScanWords)

	// Scan using our scanner until our scanning is complete
	for scanner.Scan() {
		// Print each token returned from scanner.Scan()
		// which is stored in scanner.Text(), in this
		// case, it will be word by word from string s
		fmt.Println(scanner.Text())
	}

	// Check if any scanner errors occurred and if so log them
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

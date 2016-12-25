package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	// Create an artificial input source.
	const input = "This is some awesome input, created artificially for input"

	// Create a new buffered scanner from the input source but since
	// bufio.NewScanner takes a reader and input is just plain text,
	// we use strings.NewReader(input) which returns a reader for
	// the text, which can then be used with bufio.NewScanner
	scanner := bufio.NewScanner(strings.NewReader(input))

	// Set the scanner split function to bufio.ScanWords
	scanner.Split(bufio.ScanWords)

	// Create a counter integer to increment and store the
	// amount of words through each iteration of scanner.Scan()
	count := 0

	// Scan through the reader until io.EOF or an error occurs
	for scanner.Scan() {
		// Increment the counter variable count by 1
		count++

		// Print the word at the current scanner position
		// using scanner.Text() with fmt.Println() to place
		// each word on it's own separate line
		fmt.Println(scanner.Text())
	}

	// Check if any scanner errors occurred and if so log them
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	// Print the number of words that were scanned in from the
	// input source using the bufio scanner
	fmt.Println("Amount of words:", count)
}

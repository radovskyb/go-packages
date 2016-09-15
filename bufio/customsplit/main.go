// Uses a scanner with a custom split function to validate a 32-bit
// decimal input (built by wrapping ScanWords)
package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// For this example, the result will display:
	// 1234
	// 5678
	// Lastly here will be an error describing that the last
	// token `123456789012345` is out of range. This is simply
	// because in our ParseInt call during our split function,
	// we are only allowing a maximum of 32-bits at a time for
	// each token
	const input = "1234 5678 123456789012345"

	// Create a new strings reader to read from the string input
	rdr := strings.NewReader(input)

	// Create a new scanner that scans from our string reader
	scanner := bufio.NewScanner(rdr)

	// To instead read from os.Stdin or another reader, we could
	// simply use and just comment out the above 3 lines of code.
	// This example creates a new buffered scanner from os.Stdin
	// scanner := bufio.NewScanner(os.Stdin)

	// Create a custom split function by wrapping the existing
	// ScanWords function (ScanWords is a split function
	// for a Scanner that returns each space-separated word
	// of text, with surrounding spaces deleted. It will
	// never return an empty string. The definition of space is set by
	// unicode.IsSpace. )
	split := func(data []byte, atEOF bool) (
		advance int, token []byte, err error,
	) {
		// Run scan words on our data
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			// convert our string token in to base 10
			// digits with the length 32-bits long
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		// Return our 3 return values (advance, token, err)
		return
	}

	// Set the split function for the scanning operation
	// Split sets the split function for the Scanner. If called, it must be
	// called before Scan. The default split function is ScanLines.
	scanner.Split(split)

	// Validate the input by iterating over each time the scanner
	// scans, and print the scanned line, as well as check and log
	// any errors that may occur
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	// Find and print any scanner errors using scanner.Err()
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}

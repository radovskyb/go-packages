package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a new buffered scanner to read input from os.Stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Each time input is received from os.Stdin, scan it using
	// our scanner and print out each line.
	//
	// Scan advances the Scanner to the next token, which will then be
	// available through the Bytes or Text method. It returns false when the
	// scan stops, either by reaching the end of the input or an error.
	// After Scan returns false, the Err method will return any error that
	// occurred during scanning, except that if it was io.EOF, Err
	// will return nil.
	// Scan panics if the split function returns 100 empty tokens without
	// advancing the input. This is a common error mode for scanners.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Check if there were any scanner errors and if there
	// was, print them our to os.Stderr
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
}

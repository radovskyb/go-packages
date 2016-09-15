package main

import "fmt"

func main() {
	var firstname, lastname string

	fmt.Println("Please enter your name:")

	// Scan scans text read from standard input, storing successive
	// space-separated values into successive arguments.  Newlines count
	// as space.  It returns the number of items successfully scanned.
	// If that is less than the number of arguments, err will report why.
	fmt.Scan(&firstname, &lastname)

	fmt.Printf("Hello %s %s.\n", firstname, lastname)
}

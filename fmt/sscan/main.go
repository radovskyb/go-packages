package main

import "fmt"

func main() {
	var firstname string

	// Sscan scans the argument string, storing successive space-separated
	// values into successive arguments.  Newlines count as space.  It
	// returns the number of items successfully scanned.  If that is less
	// than the number of arguments, err will report why.
	fmt.Sscan("Benjamin", &firstname)

	fmt.Printf("Hello, %s.\n", firstname)
}

package main

import (
	"fmt"
	"log"
)

func main() {
	var firstname, lastname string

	fmt.Println("Please enter your name:")

	// Scan scans text read from standard input, storing successive
	// space-separated values into successive arguments.  Newlines count
	// as space.  It returns the number of items successfully scanned.
	// If that is less than the number of arguments, err will report why.
	n, err := fmt.Scan(&firstname, &lastname)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Scanned %d words\nHello %s %s.\n", n, firstname, lastname)
}

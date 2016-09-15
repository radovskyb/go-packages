package main

import "fmt"

func main() {
	var number int

	fmt.Print("Please enter a number: ")

	// Scanf scans text read from standard input, storing successive
	// space-separated values into successive arguments as determined by
	// the format. It returns the number of items successfully scanned.
	// If that is less than the number of arguments, err will report why.
	// Newlines in the input must match newlines in the format.
	fmt.Scanf("%d", &number)

	fmt.Printf("You chose %d.\n", number)
}

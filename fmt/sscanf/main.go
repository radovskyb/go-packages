package main

import "fmt"

func main() {
	var title string
	var number int

	// Sscanf scans the argument string, storing successive space-separated
	// values into successive arguments as determined by the format.  It
	// returns the number of items successfully parsed.
	// Newlines in the input must match newlines in the format.
	fmt.Sscanf("Number: 5", "%s%d", &title, &number)

	fmt.Println(title, number)
}

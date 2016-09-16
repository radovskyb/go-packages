package main

import (
	"fmt"
	"log"
)

func main() {
	var title string
	var number int

	// Sscanf scans the argument string, storing successive space-separated
	// values into successive arguments as determined by the format.  It
	// returns the number of items successfully parsed.
	// Newlines in the input must match newlines in the format.
	_, err := fmt.Sscanf("Number: 5", "%s%d", &title, &number)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(title, number)
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var answer string

	fmt.Println("Do you like cake? (yes/no)")

	// Read one word from os.Stdin into answer
	//
	// Fscanf scans text read from r, storing successive space-separated
	// values into successive arguments as determined by the format.  It
	// returns the number of items successfully parsed.
	// Newlines in the input must match newlines in the format.
	fmt.Fscanf(os.Stdin, "%s", &answer)

	answer = strings.ToLower(answer)
	if answer == "yes" || answer[0] == byte('y') {
		fmt.Println("You like cake")
	} else {
		fmt.Println("You don't like cake")
	}
}

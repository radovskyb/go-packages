package main

import "fmt"

func main() {
	// Sprintf formats according to a format specifier and
	// returns the resulting string.
	numStr := fmt.Sprintf("Number: %d", 5)
	fmt.Println(numStr)
}

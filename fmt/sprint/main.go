package main

import "fmt"

func main() {
	// Sprint formats using the default formats for its operands and
	// returns the resulting string.
	// Spaces are added between operands when neither is a string.
	str := fmt.Sprint(1, 2, 3, " fmt.Sprint")
	fmt.Println(str)
}

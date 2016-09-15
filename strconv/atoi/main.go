package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Create a string with it's value as "10"
	v := "10"

	// Convert the string v to an integer (string "10" becomes integer 10)
	//
	// Atoi is shorthand for ParseInt(s, 10, 0).
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v\n", s, s)
		if s == 10 {
			fmt.Println("s is an integer")
		}
	}
}

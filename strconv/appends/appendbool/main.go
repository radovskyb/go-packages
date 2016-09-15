package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Create a []byte b
	b := []byte("bool: ")

	// Append a converted bool to b
	//
	// AppendBool appends "true" or "false", according to the value of b,
	// to dst and returns the extended buffer.
	b = strconv.AppendBool(b, true)

	// Print b as a string
	fmt.Println(string(b))
}

package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	t := time.Now()

	// Create a new time format to format time `t`
	format := "3:04pm"

	// Format time `t` and store the formatted []byte with the given
	// []byte below prepended to it, into variable `formatted`
	//
	// AppendFormat is like Format but appends the textual
	// representation to b and returns the extended buffer.
	formatted := t.AppendFormat([]byte("The time is: "), format)

	// Print out `formatted` as a string
	fmt.Printf("%s\n", formatted)
}

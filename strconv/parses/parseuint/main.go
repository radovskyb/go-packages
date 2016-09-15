package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "42"

	// ParseUint is like ParseInt but for unsigned numbers.
	if s, err := strconv.ParseUint(v, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	if s, err := strconv.ParseUint(v, 16, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}

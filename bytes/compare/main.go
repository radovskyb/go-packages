package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Interpret Compare's result by comparing it to zero.
	var a, b []byte

	a = []byte("a")
	b = []byte("b")

	if bytes.Compare(a, b) < 0 {
		// a less b
		fmt.Println("a < b")
	}
	if bytes.Compare(a, b) <= 0 {
		// a less or equal b
		fmt.Println("a <= b")
	}
	if bytes.Compare(a, b) > 0 {
		// a greater b
		fmt.Println("a > b")
	}
	if bytes.Compare(a, b) >= 0 {
		// a greater or equal b
		fmt.Println("a >= b")
	}

	// Prefer Equal to Compare for equality comparisons.
	if bytes.Equal(a, b) {
		// a equal b
		fmt.Println("a == b")
	}
	if !bytes.Equal(a, b) {
		// a not equal b
		fmt.Println("a != b")
	}

}

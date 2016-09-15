package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring r with 10 elements
	r := ring.New(10)

	// Populate ring r
	for i := r.Len(); i > 0; i-- {
		r.Value = i
		r = r.Prev() // Prev element to populate
	}

	// Print out ring r's element's values
	fmt.Printf("%d ", r.Value)
	reverse := r.Prev()
	for ; reverse != r; reverse = reverse.Prev() {
		fmt.Printf("%d ", reverse.Value)
	}
	fmt.Println()

	// Unlink 5 elements from ring r
	//
	// Unlink removes n % r.Len() elements from the ring r, starting
	// at r.Next(). If n % r.Len() == 0, r remains unchanged.
	// The result is the removed subring. r must not be empty.
	r.Unlink(5)

	// Print out ring r's element's values after calling r.Unlink(5)
	fmt.Printf("%d ", r.Value)
	reverse = r.Prev()
	for ; reverse != r; reverse = reverse.Prev() {
		fmt.Printf("%d ", reverse.Value)
	}
	fmt.Println()
}

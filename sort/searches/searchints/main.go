package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create a new ints slice ints
	ints := []int{5, 1, 4, 6, 3, 3, 1, 2, 9, 10}

	// Sort the ints slice ints
	sort.Ints(ints)

	fmt.Println("Sorted ints slice `ints`:", ints)

	// SearchInts searches for x in a sorted slice of ints
	// and returns the index as specified by Search.
	// The return value is the index to insert x if x is not
	// present (it could be len(a)).
	// The slice must be sorted in ascending order.
	fmt.Println(sort.SearchInts(ints, 7))
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{5, 2, 7, 1, 0, 1}

	fmt.Println("Before reverse sort:", ints)

	// Convert ints to a sort.IntSlice type and pass it
	// to sort.Reverse and then pass that to sort.Sort
	//
	// Reverse returns the reverse order for data.
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))

	fmt.Println("After reverse sort:", ints)
}

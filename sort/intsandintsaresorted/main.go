package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{9, 3, 5, 8, 9, 2, 1, 1, 0, 2, 3}

	// IntsAreSorted tests whether a slice of ints is sorted
	// in increasing order.
	if !sort.IntsAreSorted(intSlice) {
		fmt.Println("Ints are not yet sorted.")
	}
	fmt.Println("Before sort:", intSlice)

	// Ints sorts a slice of ints in increasing order.
	sort.Ints(intSlice)
	if sort.IntsAreSorted(intSlice) {
		fmt.Println("Ints are now sorted.")
	}
	fmt.Println("After sort:", intSlice)
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create a new integer slice intSlice
	intSlice := []int{1, 5, 2, 6, 9, 11, 31, 4, 6}

	// First sort the []int for search to be able to work properly
	sort.Ints(intSlice)

	// Print the now sorted intSlice
	fmt.Println("Sorted intSlice:", intSlice)

	// Create a new int to search for in `intSlice`
	x := 5

	// Use sort.Search to search for x in intSlice
	//
	// Search uses binary search to find and return the smallest index i
	// in [0, n) at which f(i) is true, assuming that on the range [0, n),
	// f(i) == true implies f(i+1) == true.  That is, Search requires that
	// f is false for some (possibly empty) prefix of the input range [0, n)
	// and then true for the (possibly empty) remainder; Search returns
	// the first true index.  If there is no such index, Search returns n.
	// (Note that the "not found" return value is not -1 as in, for instance,
	// strings.Index).
	// Search calls f(i) only for i in the range [0, n).
	//
	// A common use of Search is to find the index i for a value x in
	// a sorted, indexable data structure such as an array or slice.
	// In this case, the argument f, typically a closure, captures the value
	// to be searched for, and how the data structure is indexed and
	// ordered.
	//
	// For instance, given a slice data sorted in ascending order,
	// the call Search(len(data), func(i int) bool { return data[i] >= 23 })
	// returns the smallest index i such that data[i] >= 23.  If the caller
	// wants to find whether 23 is in the slice, it must test data[i] == 23
	// separately.
	//
	// Searching data sorted in descending order would use the <=
	// operator instead of the >= operator.
	i := sort.Search(len(intSlice), func(i int) bool {
		return intSlice[i] >= x
	})

	// Check if x was found in intSlice
	if i < len(intSlice) && intSlice[i] == x {
		fmt.Println(x, "is present in intSlice at position", i)
	} else {
		fmt.Println(x, "is not present in intSlice")
		fmt.Println("current position:", i)
	}
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create a new float64 slice f64s
	f64s := []float64{4.23, 3.21, 1.55, 1.01, 2.101}

	// Sort the float64 slice f64s
	sort.Float64s(f64s)

	fmt.Println("Sorted float64 slice `f64s`:", f64s)

	// SearchFloat64s searches for x in a sorted slice of
	// float64s and returns the index as specified by Search.
	// The return value is the index to insert x if x is not
	// present (it could be len(a)).
	// The slice must be sorted in ascending order.
	fmt.Println(sort.SearchFloat64s(f64s, 5))
}

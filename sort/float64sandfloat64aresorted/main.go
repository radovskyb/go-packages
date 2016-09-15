package main

import (
	"fmt"
	"sort"
)

func main() {
	float64Slice := []float64{1.53123, 6.123123, 6.55555, 1.11111}

	// Float64sAreSorted tests whether a slice of float64s is
	// sorted in increasing order.
	if !sort.Float64sAreSorted(float64Slice) {
		fmt.Println("Floats are not yet sorted.")
	}
	fmt.Println("Before sort:", float64Slice)

	// Float64s sorts a slice of float64s in increasing order.
	sort.Float64s(float64Slice)
	if sort.Float64sAreSorted(float64Slice) {
		fmt.Println("Floats are now sorted.")
	}
	fmt.Println("After sort:", float64Slice)
}

package main

import "fmt"

func main() {
	// Create 2 integer slices
	slice_one := []int{1, 2, 3}
	slice_two := []int{4, 5, 6}

	// Print both slice's contents before appending
	// slice_two onto slice_one
	PrintSlice("slice_one before appending slice_two", slice_one)
	PrintSlice("slice_two before append to slice_one", slice_two)

	fmt.Println()

	// Store the results of appending slice_two onto
	// slice_one, in the slice_one slice
	slice_one = append(slice_one, slice_two...)

	// Print both slice's contents after appending slice_two
	// onto slice_one. slice_one is now [1, 2, 3, 4, 5, 6]
	PrintSlice("slice_one after appending slice_two", slice_one)
	PrintSlice("slice_two after append to slice_one", slice_two)
}

func PrintSlice(name string, slice []int) {
	// Print the slice name before showing its contents
	fmt.Printf("%s: ", name)

	// Iterate over the range of the slice and print
	// each value followed by one space character
	for _, v := range slice {
		fmt.Print(v, " ")
	}

	// Print one empy line after the slice's contents are printed
	fmt.Println()
}

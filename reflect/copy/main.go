package main

import (
	"fmt"
	"reflect"
)

func main() {
	oldSlice := []string{"abc", "def"}
	newSlice := []string{"rst", "uvw", "xyz"}

	// `rst` and `uvw` elements from newSlice will be overwritten by
	// `abc` and `def` when copying oldSlice into newSlice
	//
	// Copy copies the contents of src into dst until either
	// dst has been filled or src has been exhausted.
	// It returns the number of elements copied.
	// Dst and src each must have kind Slice or Array, and
	// dst and src must have the same element type.
	var n = reflect.Copy(
		reflect.ValueOf(newSlice),
		reflect.ValueOf(oldSlice),
	)

	fmt.Println("Copied", n, "elements")
	fmt.Printf("newSlice's elements are now %v\n", newSlice)
}

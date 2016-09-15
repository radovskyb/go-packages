package main

import (
	"fmt"
	"reflect"
)

func main() {
	sliceOne := []string{"abc", "def"}
	sliceTwo := []string{"abc", "def"}

	// Check if sliceOne and sliceTwo are equal using reflect.DeepEqual.
	// Two slices can't be compared using the == symbol, for example:
	// `if sliceOne == sliceTwo {
	// 	fmt.Println("Both slices are equal")
	// }`
	// The above is not valid go code so instead reflect.DeepEqual is used.
	//
	// DeepEqual tests for deep equality. It uses normal == equality where
	// possible but will scan elements of arrays, slices, maps,
	// and fields of structs. In maps, keys are compared with == but
	// elements use deep equality. DeepEqual correctly handles recursive
	// types. Functions are equal only if they are both nil.
	// An empty slice is not equal to a nil slice.
	if reflect.DeepEqual(sliceOne, sliceTwo) {
		fmt.Println("Both slices are equal")
	}
}

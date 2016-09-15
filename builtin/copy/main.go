// The copy built-in function copies elements from a source
// slice into a destination slice. (As a special case, it also
// will copy bytes from a string to a slice of bytes.) The
// source and destination may overlap. Copy returns the number of
// elements copied, which will be the minimum of len(src) and len(dst).
// func copy(dst, src []Type) int <- Function signature
package main

import "fmt"

func main() {
	// Write out `Slice One` header
	fmt.Println("Slice One:")

	// Create a slice, slice_one, that can hold 1 string
	slice_one := make([]string, 1)

	// Create a string, str, that contains the text `Hello, World!`
	str := "Hello, World!"

	// Print out slice_one before the copy
	fmt.Printf("slice_one before copy: %v\n", slice_one)

	// Copy the string from str into slice_one and store how
	// many elements were copied in the integer n
	n := copy(slice_one, []string{str})

	// Print out how many elements were copied into slice_one which
	// can only be one in this case since we are passing a single string
	fmt.Printf("\nCopied %d element/s into the slice_one\n", n)

	// Print out slice_one after the copy
	fmt.Printf("slice_one after copy: %v\n", slice_one)

	// Write out `Slice Two` header
	fmt.Println("\nSlice Two:")

	// Create a slice, slice_two, that can hold 2 strings
	slice_two := make([]string, 2)

	// Create a string slice, strings, that contains two strings
	strings := []string{"Hello, World!", "How are you?"}

	// Print out slice_two before the copy
	fmt.Printf("slice_two before copy: %v\n", slice_two)

	// Copy the strings from string slice strings into slice_two
	// and store how many elements were copied in the integer n
	n = copy(slice_two, strings)

	// Print out how many elements were copied into slice_two
	fmt.Printf("\nCopied %d element/s into the slice_two\n", n)

	// Print out slice_two after the copy
	fmt.Printf("slice_two after copy: %v", slice_two)
}

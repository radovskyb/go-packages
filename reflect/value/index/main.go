package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create a new string slice s
	s := []string{"abc", "def", "ghi"}

	// Create a reflect.Value object type v, from string slice s
	v := reflect.ValueOf(s)

	// Print v's first element using v.Index(0)
	//
	// Index returns v's i'th element.
	// It panics if v's Kind is not Array, Slice, or String or
	// i is out of range.
	fmt.Println(v.Index(0))

	// Iterate over v and print out all of it's elements
	for i := 0; i < v.Len(); i++ {
		fmt.Printf("%v ", v.Index(i))
	}
}

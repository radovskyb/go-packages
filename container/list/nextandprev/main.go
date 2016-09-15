package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Create a new list
	l := list.New()

	// Push 2 elements onto the list
	l.PushBack(1)
	l.PushBack(2)

	// Get the element at the front of the list
	eFront := l.Front()

	// Print eFront's value
	fmt.Println(eFront.Value)

	// Get the next element after eFront
	//
	// Next returns the next list element or nil.
	next := eFront.Next()

	// Print out next's value
	fmt.Println(next.Value)

	// Now print out next's previous element's value just
	// to show that it will print out eFront's value again
	//
	// Prev returns the previous list element or nil.
	fmt.Println(next.Prev().Value)
}

package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Create a new list
	l := list.New()

	// Push 5 elements onto the list
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	// Print the length of list l
	//
	// Len returns the number of elements of list l.
	// The complexity is O(1).
	fmt.Println(l.Len()) // 5

	// Clear list l by using l.Init()
	//
	// Init initializes or clears list l.
	l.Init()

	// Print list l's length again after calling l.Init()
	fmt.Println(l.Len()) // 0
}

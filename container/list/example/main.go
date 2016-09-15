package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Create a new list
	//
	// New returns an initialized list.
	l := list.New()

	// Push an integer 4 to the back of the list
	//
	// PushBack inserts a new element e with value v
	// at the back of list l and returns e.
	e4 := l.PushBack(4)

	// Push an integer 1 to the front of the list
	//
	// PushFront inserts a new element e with value v at
	// the front of list l and returns e.
	e1 := l.PushFront(1)

	// Insert a new integer 3 before element e4
	//
	// InsertBefore inserts a new element e with value v
	// immediately before mark and returns e.
	// If mark is not an element of l, the list is not modified.
	l.InsertBefore(3, e4)

	// Insert a new integer 2 after element e1
	//
	// InsertAfter inserts a new element e with value v
	// immediately after mark and returns e.
	// If mark is not an element of l, the list is not modified.
	l.InsertAfter(2, e1)

	// Iterate through the list and print it's contents
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}

	// The above loop prints 1 2 3 4
}

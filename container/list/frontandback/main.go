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

	// Get the first element in the list
	//
	// Front returns the first element of list l or nil.
	eFront := l.Front()

	// Get the last element in the list
	//
	// Back returns the last element of list l or nil.
	eBack := l.Back()

	// Print eFront's value
	fmt.Println(eFront.Value) // 1

	// Print eBack's value
	fmt.Println(eBack.Value) // 5
}

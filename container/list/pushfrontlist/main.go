package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Create a new list l1
	l1 := list.New()

	// Push 2 elements on to the back of list l1
	l1.PushBack(1)
	l1.PushBack(2)

	// Print list l1
	printList("l1", l1)

	// Now create a second list l2
	l2 := list.New()

	// Push 3 elements on to the back of list l2
	l2.PushBack(3)
	l2.PushBack(4)
	l2.PushBack(5)

	// Print list l2
	printList("l2", l2)

	// Now push a copy of list l2, onto the front of list l1
	//
	// PushFrontList inserts a copy of an other list at the front of list l.
	// The lists l and other may be the same.
	l1.PushFrontList(l2)

	// Print l1 after pushing l2 onto the front of it
	printList("l1 after pushing l2 onto the front of it", l1)
}

// printList takes a list's name and a list and print's out first
// the name of the list, and then all of the list's elements' values
func printList(name string, list *list.List) {
	fmt.Print(name + ": ")
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println()
}

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

	// Print list l
	printList(l)

	// Get the element at the front of the list
	e1 := l.Front()

	// Move e1 after the second element in the list
	//
	// MoveAfter moves element e to its new position after mark.
	// If e or mark is not an element of l, or e == mark,
	// the list is not modified.
	l.MoveAfter(e1, e1.Next())

	// Print the list again after moving e1 after
	// the second element in the list
	printList(l)

	// Now move e1 which is now technically element 2 in the list,
	// back to the first position, by placing it back in front of
	// what is now the first element in the list
	//
	// MoveBefore moves element e to its new position before mark.
	// If e or mark is not an element of l, or e == mark, the
	// list is not modified.
	l.MoveBefore(e1, e1.Prev())

	// Print the list again after moving the element back
	// and it will print the same way it first did
	printList(l)

	// Now take the first element again, but this time, move
	// it to the back of the list
	//
	// MoveToBack moves element e to the back of list l.
	// If e is not an element of l, the list is not modified.
	l.MoveToBack(e1)

	// Print the list to see e1 at the back of the list
	printList(l)

	// Lastly, move the last element which is element e1, back
	// to the front of the list
	//
	// MoveToFront moves element e to the front of list l.
	// If e is not an element of l, the list is not modified.
	l.MoveToFront(e1)

	// And finally, print the list again to see the list back
	// to normal once again
	printList(l)
}

// Print list takes a list as a parameter, iterates over
// it and prints out each element's value
func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println()
}

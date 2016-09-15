package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Create a new list l
	l := list.New()

	// Push 3 items onto list l
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	// Print list l
	printList(l)

	// Remove the front element from the list and store it's
	// value in v
	//
	// Remove removes e from l if e is an element of list l.
	// It returns the element value e.Value.
	v := l.Remove(l.Front())

	// Print the list after removing the first element
	printList(l)

	// Print the value of the element that was removed/popped
	// from the list
	fmt.Println(v)

	// Now put it back on the fron of the list
	l.PushFront(v)

	// Print the list one last time with the elements back in tact
	printList(l)
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println()
}

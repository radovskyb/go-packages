package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of 5 elements
	//
	// New creates a ring of n elements.
	r := ring.New(5)

	// Print the length of ring r
	//
	// Len computes the number of elements in ring r.
	// It executes in time proportional to the number of elements.
	fmt.Println(r.Len()) // 5

	fmt.Println()

	i := 0
	f1 := func(_ interface{}) {
		// Use move to move to the ring element at position i
		// and store i's value in each position
		//
		// Move moves n % r.Len() elements backward (n < 0)
		// or forward (n >= 0) in the ring and returns that ring
		// element. r must not be empty.
		r.Move(i).Value = i

		// Print out the rings value at the i'th position
		fmt.Printf("%d\n", r.Move(i).Value)

		// Increment i
		i++
	}

	// Call the f function with ring r. The f function simply gives
	// each element in the ring a value according to it's position in
	// the ring
	//
	// Do calls function f on each element of the ring, in forward order.
	// The behavior of Do is undefined if f changes *r.
	r.Do(f1)

	fmt.Println()

	// Print out all of the elements in the ring using a for loop
	for i := 0; i < r.Len(); i++ {
		fmt.Println(r.Move(i).Value)
	}

	fmt.Println()

	// Create a new function to be called with r.Do which just prints
	// out each element's value in the ring
	f2 := func(x interface{}) {
		fmt.Println(x)
	}

	// Call r.Do with f2 as it's function parameter
	r.Do(f2) // Prints out all ring elements

	fmt.Println()

	// Print out the last element in the ring by calling r.Prev()
	// from the ring at it's current starting position
	//
	// Prev returns the previous ring element. r must not be empty.
	fmt.Println(r.Prev().Value) // 4

	fmt.Println()

	// From this same starting position, call r.Next() which will
	// print the element at the next position after the starting
	// position
	//
	// Next returns the next ring element. r must not be empty.
	fmt.Println(r.Next().Value) // 1

	fmt.Println()

	// Create a loop which takes each element and doubles it's value
	for i := 0; i < r.Len(); i++ {
		// Take the current element's value and double it
		r.Value = r.Value.(int) * 2

		// Move r to the next element in the ring
		r = r.Next()
	}

	// Call r.Do with function f2 again to print out each element in
	// the ring again, however this time, each element will be doubled
	r.Do(f2) // Prints each of the rings elements doubled from earlier
}

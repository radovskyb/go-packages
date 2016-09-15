package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a first ring r1 with 5 elements
	r1 := ring.New(5)

	// Create a function to display a ring's element values
	// to be used with func (r *Ring) Do(f func(interface{}))
	f := func(x interface{}) {
		fmt.Printf("%d ", x)
	}

	// Populate ring r1
	for i := 1; i < r1.Len()+1; i++ {
		r1.Value = i
		r1 = r1.Next()
	}

	// Print r1's element's values
	fmt.Println("Ring r1")
	r1.Do(f)

	fmt.Println()

	// Create a second ring r2 with 5 elements
	r2 := ring.New(5)

	// Populate ring r2
	for i := 1; i < r2.Len()+1; i++ {
		r2.Value = i * 10
		r2 = r2.Next()
	}

	// Print r2's element's values
	fmt.Println("\nRing r2")
	r2.Do(f)

	fmt.Println()

	// Link ring r1 with ring r2
	//
	// Link connects ring r with ring s such that r.Next()
	// becomes s and returns the original value for r.Next().
	// r must not be empty.
	//
	// If r and s point to the same ring, linking
	// them removes the elements between r and s from the ring.
	// The removed elements form a subring and the result is a
	// reference to that subring (if no elements were removed,
	// the result is still the original value for r.Next(),
	// and not nil).
	//
	// If r and s point to different rings, linking
	// them creates a single ring with the elements of s inserted
	// after r. The result points to the element following the
	// last element of s after insertion.
	r1.Link(r2)

	fmt.Println()

	// Now print ring r1 after linking it with ring r2
	fmt.Println("Ring r1 after linking it with r2")
	r1.Do(f)

	fmt.Println()

	fmt.Println("\nRing r2 after being linked with r1")
	r2.Do(f)

	fmt.Println()
	fmt.Println()

	// ring r1's next value is now the first value of ring r2
	fmt.Println(
		"r1's next value is now the first value of r2:",
		r1.Next().Value,
	)

	// ring r2's prev value is now the first value of ring r1
	fmt.Println(
		"r2's prev value is now the first value of r1:",
		r2.Prev().Value,
	)
}

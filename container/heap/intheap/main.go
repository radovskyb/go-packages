// This example demonstrates an integer heap built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// The following two IntHeap methods Push and Pop use pointer receivers
// because they modify the slice's length, not just its contents.
func (h *IntHeap) Push(x interface{}) {
	// Append x as an int into the IntHeap slice
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	// Store the heap as it is now
	old := *h

	// Get the length of the heap as it is now
	n := len(old)

	// Create x which is set to the last element from the current heap
	x := old[n-1]

	// Now set the heap to the old heap minus the last element therefore
	// removing the last element (popping it) from the heap
	*h = old[0 : n-1]

	// Return x, the last element from the heap which was popped off
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	// Create a new IntHeap h
	//
	// Heap h is unsorted but will be popped off in order (priority)
	h := &IntHeap{1, 10, 5, 25, 0}

	// Initialize h as a heap type using heap.Init(h)
	//
	// A heap must be initialized before any of the heap operations
	// can be used. Init is idempotent with respect to the heap invariants
	// and may be called whenever the heap invariants may have been
	// invalidated.
	// Its complexity is O(n) where n = h.Len().
	heap.Init(h)

	// Push an integer 15 on to the heap h
	//
	// Push pushes the element x onto the heap. The complexity is
	// O(log(n)) where n = h.Len().
	heap.Push(h, 15)

	fmt.Printf("minimum: %d\n", (*h)[0])

	for h.Len() > 0 {
		// Pop off each element one by one until the loop finishes
		// when there are no more elements on the heap
		//
		// Pop removes the minimum element (according to Less) from the heap
		// and returns it. The complexity is O(log(n)) where n = h.Len().
		// It is equivalent to Remove(h, 0).
		fmt.Printf("%d ", heap.Pop(h))
	}
}

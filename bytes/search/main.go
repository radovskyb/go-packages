package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	// Binary search to find a matching byte slice.
	var needle []byte
	var haystack [5][]byte

	needle = []byte("c")

	fmt.Println("Needle: ", string(needle))

	haystack[0] = []byte("c")
	haystack[1] = []byte("a")
	haystack[2] = []byte("b")
	haystack[3] = []byte("d")
	haystack[4] = []byte("e")

	// Sort the haystack byte arrays based on bytes.Compare
	for k, v := range haystack {
		for k1, v1 := range haystack {
			if bytes.Compare(v, v1) < 0 {
				haystack[k], haystack[k1] = haystack[k1], haystack[k]
			}
		}
	}

	// Print out the sorted haystack's values
	fmt.Print("Sorted haystack: ")
	for _, v := range haystack {
		fmt.Print(string(v), " ")
	}
	fmt.Println()

	// Now that the haystack is sorted, we can run sort.Search
	// to find out if our needle is in our haystack
	i := sort.Search(len(haystack), func(i int) bool {
		// Return haystack[i] >= needle.
		return bytes.Compare(haystack[i], needle) >= 0
	})

	// If the needle is found in the haystack, print out what
	// position the needle was found at, otherwise print not found
	if i < len(haystack) && bytes.Equal(haystack[i], needle) {
		fmt.Println("Found at pos:", i)
	} else {
		fmt.Println("Not found")
	}
}

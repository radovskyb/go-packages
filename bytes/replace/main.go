package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Create a new byte slice b
	b := []byte{'a', 'b', 'c', 'a', 'b', 'c'}

	// Replace any `a` bytes with `z` bytes, a maximum of
	// two times, from the `b` byte slice, and storing the
	// results in byte slice `b1`.
	//
	// Replace returns a copy of the slice s with the first n
	// non-overlapping instances of old replaced by new. If old
	// is empty, it matches at the beginning of the slice and after
	// each UTF-8 sequence, yielding up to k+1 replacements for a
	// k-rune slice. If n < 0, there is no limit on the number
	// of replacements.
	b1 := bytes.Replace(b, []byte{'a'}, []byte{'z'}, 2)

	// Print out the stringified byte slice b1.
	fmt.Println(string(b1))
}

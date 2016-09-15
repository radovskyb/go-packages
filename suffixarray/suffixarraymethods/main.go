package main

import (
	"fmt"
	"index/suffixarray"
	"regexp"
)

func main() {
	// Create a new string containing some string data
	data := "La dee dah dee dah dee dah de dee dah dee dah dee da! :)"

	// Create a new Index object for data
	//
	// New creates a new Index for data.
	// Index creation time is O(N*log(N)) for N = len(data).
	//
	// Index implements a suffix array for fast substring search.
	index := suffixarray.New([]byte(data))

	// Print out the index's bytes as a string
	//
	// Bytes returns the data over which the index was created.
	// It must not be modified.
	fmt.Println(string(index.Bytes()))

	// Return all indices where the bytes `dee` occur in `data`
	//
	// Lookup returns an unsorted list of at most n indices where the byte string s
	// occurs in the indexed data. If n < 0, all occurrences are returned.
	// The result is nil if s is empty, s is not found, or n == 0.
	// Lookup time is O(log(N)*len(s) + len(result)) where N is the
	// size of the indexed data.
	offsets := index.Lookup([]byte("dee"), -1)

	// Print out all of the offsets
	fmt.Println(offsets)

	// Print out at most 3 of the indices where `dee` occurs in `data`
	offsets = index.Lookup([]byte("dee"), 3)

	// Print out all of the offsets
	fmt.Println(offsets)

	// Create a regexp to search for inside of `data`
	// Try to find anything that contains de followed by any other characters
	match := regexp.MustCompile("dee")

	// Use the above regex to find all matches in the data
	//
	// FindAllIndex returns a sorted list of non-overlapping matches of the
	// regular expression r, where a match is a pair of indices specifying
	// the matched slice of x.Bytes(). If n < 0, all matches are returned
	// in successive order. Otherwise, at most n matches are returned and
	// they may not be successive. The result is nil if there are no matches,
	// or if n == 0.
	result := index.FindAllIndex(match, -1)

	// Print out all `dee`'s and their starting positions
	for _, val := range result {
		start, end := val[0], val[1]
		fmt.Println(start, data[start:end])
	}
}

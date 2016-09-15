package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create a new strings slice strings
	strings := []string{"zzz", "xxx", "bbb", "ccc", "aaa"}

	// Sort the strings slice strings
	sort.Strings(strings)

	fmt.Println("Sorted strings slice `strings`:", strings)

	// SearchStrings searches for x in a sorted slice of strings
	// and returns the index as specified by Search. The return value is
	// the index to insert x if x is not present (it could be len(a)).
	// The slice must be sorted in ascending order.
	fmt.Println(sort.SearchStrings(strings, "yyy"))
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	stringsSlice := []string{"ccc", "bbb", "aaa", "zzz", "aaa"}

	// StringsAreSorted tests whether a slice of strings is
	// sorted in increasing order.
	if !sort.StringsAreSorted(stringsSlice) {
		fmt.Println("Strings are not yet sorted.")
	}
	fmt.Println("Before sort:", stringsSlice)

	// Strings sorts a slice of strings in increasing order.
	sort.Strings(stringsSlice)
	if sort.StringsAreSorted(stringsSlice) {
		fmt.Println("Strings are now sorted.")
	}
	fmt.Println("After sort:", stringsSlice)
}

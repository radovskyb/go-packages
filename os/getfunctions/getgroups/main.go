package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the caller user's groups
	//
	// Getgroups returns a list of the numeric ids of groups that
	// the caller belongs to.
	groups, _ := os.Getgroups()

	// Iterate over and print each group
	for _, group := range groups {
		fmt.Println(group)
	}
}

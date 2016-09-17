package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Get the caller user's groups
	//
	// Getgroups returns a list of the numeric ids of groups that
	// the caller belongs to.
	groups, err := os.Getgroups()
	if err != nil {
		log.Fatalln(err)
	}

	// Iterate over and print each group
	for _, group := range groups {
		fmt.Println(group)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Get a rooted path name corresponding to the current directory
	//
	// Getwd returns a rooted path name corresponding to the current
	// directory. If the current directory can be reached via multiple
	// paths (due to symbolic links), Getwd may return any one of them.
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	// Print the rooted path of the current directory
	fmt.Println(wd)
}

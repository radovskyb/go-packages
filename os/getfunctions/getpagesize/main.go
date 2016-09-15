package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the current underlying systems memory page size
	//
	// Getpagesize returns the underlying system's memory page size.
	pagesize := os.Getpagesize()

	// Print out the page size integer
	fmt.Println(pagesize)
}

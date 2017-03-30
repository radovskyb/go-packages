package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Stat("tmp")
	if err != nil {
		// Check if the directory does not yet exist.
		//
		// IsNotExist returns a boolean indicating whether the error is known to
		// report that a file or directory does not exist. It is satisfied by
		// ErrNotExist as well as some syscall errors.
		if os.IsNotExist(err) {
			// fmt.Println(os.IsExists(err) will print false since we know the
			// tmp folder doesn't exist.
			fmt.Printf("%t: tmp folder doesn't exist.\n", os.IsExist(err))
		} else {
			// It must be a different error since tmp exists.
			log.Fatalln(err)
		}
	}
}

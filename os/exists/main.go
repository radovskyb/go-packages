package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Mkdir("tmp", os.ModePerm)
	if err != nil {
		// Check if the directory already exists.
		//
		// IsExist returns a boolean indicating whether the error is known to report
		// that a file or directory already exists. It is satisfied by ErrExist as
		// well as some syscall errors.
		if os.IsExist(err) {
			// fmt.Println(os.IsExists(err) will print true since we know the
			// tmp file already exists.
			fmt.Printf("%t: tmp file already exists.\n", os.IsExist(err))
		} else {
			// It must be a different error since tmp doesn't exist.
			log.Fatalln(err)
		}
	}
}

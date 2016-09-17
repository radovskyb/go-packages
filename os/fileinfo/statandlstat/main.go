package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Retrieve a file info describing file.txt
	//
	// Lstat returns a FileInfo describing the named file. If the
	// file is a symbolic link, the returned FileInfo describes the
	// symbolic link. Lstat makes no attempt to follow the link. If
	// there is an error, it will be of type *PathError.
	fi, err := os.Lstat("file.txt")

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print the returned info's name (file.txt)
	fmt.Println(fi.Name())

	// Now get the file info describing file.txt using os.Stat
	//
	// Stat returns a FileInfo describing the named file. If there is an error, it will be of type *PathError.
	fi, err = os.Stat("file.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Print the returned info's size of file.txt
	fmt.Println(fi.Size())
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("file.txt")

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
		if err := os.Remove("file.txt"); err != nil {
			log.Fatalln(err)
		}
	}()

	fileInfo, err := f.Stat()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print whether the file is a directory
	//
	// IsDir reports whether m describes a directory. That is, it
	// tests for the ModeDir bit being set in m.
	fmt.Println(fileInfo.Mode().IsDir())

	// Print whether the file is a regular file
	//
	// IsRegular reports whether m describes a regular file. That is,
	// it tests that no mode type bits are set.
	fmt.Println(fileInfo.Mode().IsRegular())

	// Print the files permissions
	//
	// Perm returns the Unix permission bits in m.
	fmt.Println(fileInfo.Mode().Perm())

	// Print the file mode string
	fmt.Println(fileInfo.Mode().String())
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Get the file info for file1.txt and file2.txt
	f1, err := os.Stat("file1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	f2, err := os.Stat("file1.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Check if file1.txt and file2.txt describe the same file
	//
	// SameFile reports whether fi1 and fi2 describe the same file.
	// For example, on Unix this means that the device and inode fields
	// of the two underlying structures are identical; on other systems
	// the decision may be based on the path names. SameFile only applies to
	// results returned by this package's Stat. It returns false in other cases.
	same := os.SameFile(f1, f2)

	// Print if file1.txt and file2.txt describe the same file
	fmt.Println(same) // true

	// Now get file3.txt's info
	f3, err := os.Stat("file3.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// And now check if file1.txt and file3.txt describe the same file
	same = os.SameFile(f1, f3)

	// Now print if file1.txt and file3.txt describe the same file
	fmt.Println(same) // false
}

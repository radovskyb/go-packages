package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Create a new file file.txt
	f, err := os.Create("file.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Create a defer function to run some code after main finishes execution,
	// mainly to close and delete the file file.txt when we are done with it
	defer func() {
		// Close the file handle
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}

		// Remove the file
		if err := os.Remove("file.txt"); err != nil {
			log.Fatalln(err)
		}
	}()

	// Write a byte slice to file.txt
	_, err = f.Write([]byte("Hello, World!\n"))
	if err != nil {
		log.Fatalln(err)
	}

	// Get file.txt's file info and store it as finfo
	finfo, err := f.Stat()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Write out how many bytes file.txt contains using finfo.Size()
	fmt.Printf("file.txt is %d bytes long\n", finfo.Size()) // 14 bytes

	// Read all of file.txt's contents and store it in contents byte slice
	contents, err := ioutil.ReadFile("file.txt")

	// Check if there were any errors reading file.txt's contents
	if err != nil {
		log.Fatalln(err)
	}

	// Print out file.txt's contents as a string
	fmt.Println(string(contents))

	// Now truncate the file's size by using os.Truncate, to only 10 bytes
	// long. Therefore it will truncate the last 4 bytes of file.txt's contents
	//
	// Truncate changes the size of the named file. If the file is a
	// symbolic link, it changes the size of the link's target.
	// If there is an error, it will be of type *PathError.
	err = os.Truncate("file.txt", 10)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Once again get file.txt's file info and store it as finfo
	finfo, err = f.Stat()

	// And once again check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Once again  write out how many bytes file.txt contains after
	// using os.Truncate
	fmt.Printf("file.txt is now %d bytes long\n", finfo.Size()) // 10 bytes

	// Read all of file.txt's contents and store it in contents byte slice
	contents, err = ioutil.ReadFile("file.txt")

	// Again check if there were any errors reading file.txt's contents
	if err != nil {
		log.Fatalln(err)
	}

	// Lastly print out file.txt's contents as a string again to show
	// what it looks like after using os.Truncate on file.txt
	fmt.Println(string(contents))
}

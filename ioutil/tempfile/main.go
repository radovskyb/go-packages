package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Create a new temporary directory to store a temporary file
	tmpDir, err := ioutil.TempDir("./", "tmp")
	if err != nil {
		log.Fatalln(err)
	}

	// Now create a new temporary file inside of the temporary directory
	//
	// TempFile creates a new temporary file in the directory dir with
	// a name beginning with prefix, opens the file for reading and writing,
	// and returns the resulting *os.File. If dir is the empty string,
	// TempFile uses the default directory for temporary files
	// (see os.TempDir). Multiple programs calling TempFile simultaneously
	// will not choose the same file. The caller can use f.Name() to find
	// the pathname of the file. It is the caller's responsibility to
	// remove the file when no longer needed.
	tmpFile, err := ioutil.TempFile(tmpDir, "tmp")
	if err != nil {
		log.Fatalln(err)
	}

	// Print out a created file statement
	fmt.Printf("Created temporary file %v in temporary directory %v.\n",
		tmpFile.Name(), tmpDir)

	// Close the file
	if err := tmpFile.Close(); err != nil {
		log.Fatalln(err)
	}

	// Print a deleting file statement
	fmt.Println("Deleting temporary file:", tmpFile.Name())

	// Remove the file
	if err := os.Remove(tmpFile.Name()); err != nil {
		log.Fatalln(err)
	}

	// Print a deleting directory statement
	fmt.Println("Deleting temporary directory:", tmpDir)

	// Remove the directory
	if err := os.Remove(tmpDir); err != nil {
		log.Fatalln(err)
	}
}

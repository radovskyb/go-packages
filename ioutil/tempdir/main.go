package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Create a new temporary directory
	//
	// TempDir creates a new temporary directory in the directory dir
	// with a name beginning with prefix and returns the path of the new
	// directory. If dir is the empty string, TempDir uses the default
	// directory for temporary files (see os.TempDir). Multiple programs
	// calling TempDir simultaneously will not choose the same directory.
	// It is the caller's responsibility to remove the directory when no
	// longer needed.
	tmpDirName, err := ioutil.TempDir("./", "tmp")

	// Check if there were any errors and if so log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print the tmp dir's name
	fmt.Printf("%s \n", tmpDirName)

	// Now remove the temporary directory now that we are done with it
	if err := os.Remove(tmpDirName); err != nil {
		log.Fatalln(err)
	}
}

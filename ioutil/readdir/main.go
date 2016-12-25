package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Read the current directory and return a list of sorted directory entries
	//
	// ReadDir reads the directory named by dirname and returns a list
	// of sorted directory entries.
	dirs, err := ioutil.ReadDir("./")

	// Check if there were any errors
	if err != nil {
		// Log any errors that occurred
		log.Fatalln(err)
	}

	// Iterate over the directories info's returned from from ioutil.ReadDir
	for _, fileInfo := range dirs {
		// Print out each of the directories names
		fmt.Println(fileInfo.Name())
	}
}

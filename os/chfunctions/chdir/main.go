package main

import (
	"io/ioutil"
	"os"
)

func main() {
	// Change to the directory go src
	//
	// Chdir changes the current working directory to the named directory.
	// If there is an error, it will be of type *PathError.
	os.Chdir("/Users/Benjamin/Workspace/go/src")

	// Write a file called file.txt to the directory go src
	// and write a byte slice to it
	ioutil.WriteFile("file.txt", []byte("Hello, World!"), 0644)
}

package main

import (
	"fmt"
	"syscall"
)

func main() {
	// Getenv retrieves the value of the environment variable named by the key.
	// It returns the value, which will be empty if the variable is not present.
	//
	// If found is false, then no env var was found for the key.
	goPath, found := syscall.Getenv("GOPATH")
	if !found {
		fmt.Println("env var for key not found")
		return
	}
	fmt.Println(goPath)
}

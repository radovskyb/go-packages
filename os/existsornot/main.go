package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("tmp")

	// Check if the directory does not yet exist
	//
	// IsNotExist returns a boolean indicating whether the error is known to
	// report that a file or directory does not exist. It is satisfied by
	// ErrNotExist as well as some syscall errors.
	notExists := os.IsNotExist(err) // true

	// Print out the true or false result of os.IsNotExist
	fmt.Println(notExists) // true since tmp doesnt yet exist

	// Try to create a new directory and see if there is any error returned
	err = os.Mkdir("tmp", 0755)

	// Check if the current directory exists using os.IsExist based on
	// the error returned above
	//
	// IsExist returns a boolean indicating whether the error is known to
	// report that a file or directory already exists. It is satisfied
	// by ErrExist as well as some syscall errors.
	exists := os.IsExist(err)

	// Print out the true or false result of os.IsExist
	fmt.Println(exists) // false since tmp folder was created above

	// Once finished, remove the directory tmp
	os.Remove("tmp")
}

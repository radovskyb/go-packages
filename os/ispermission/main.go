package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the file named file and try to open it for reading and writing
	file, err := os.OpenFile("file.txt", os.O_WRONLY, 0644)

	// There should be an error here since if file file only has read
	// permissions and the above line is trying to open the file for writing.
	if err != nil {
		// See if the error above was related to permissions errors such as
		// file not being writable
		//
		// IsPermission returns a boolean indicating whether the error is
		// known to report that permission is denied. It is satisfied by
		// ErrPermission as well as some syscall errors.
		if os.IsPermission(err) {
			fmt.Println("Permissions error")
		} else {
			log.Fatalln(err)
		}
	}

	// Close the file handle
	if err := file.Close(); err != nil {
		log.Fatalln(err)
	}
}

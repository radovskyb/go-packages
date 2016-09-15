package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the file named file and try to open it for reading and writing
	file, err := os.OpenFile("file", os.O_RDWR|os.O_APPEND, 0755)

	// defer to close the file handle
	defer file.Close()

	// There should be an error here since the file file only has read
	// permissions and the above line is trying to open it for writing to.
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

}

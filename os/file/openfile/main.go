package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Create and then open the file file.txt for reading and writing
	//
	// OpenFile is the generalized open call; most users will use Open
	// or Create instead. It opens the named file with specified flag
	// (O_RDONLY etc.) and perm, (0666 etc.) if applicable. If successful,
	// methods on the returned File can be used for I/O. If there is
	// an error, it will be of type *PathError.
	f, err := os.OpenFile("file.txt", os.O_CREATE|os.O_RDWR, 0755)

	// Defer to close the file and remove it
	defer func() {
		f.Close()
		err := os.Remove("file.txt")

		// Check if there were any errors and if so, log them
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Write the contents from the file handle for file.txt to os.Stdout
	io.Copy(os.Stdout, f)

	// Create a new bytes reader that reads from a byte slice
	br := bytes.NewReader([]byte("What's up?"))

	// Write from the bytes reader br to the file
	_, err = br.WriteTo(f)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Now read the full contents of file.txt into byte slice contents
	contents, err := ioutil.ReadFile("file.txt")

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print file.txt's contents
	fmt.Println(string(contents)) // What's up?
}

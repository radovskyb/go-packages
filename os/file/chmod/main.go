package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a new file with the permissions 0644
	f, err := os.OpenFile("file.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	// When main finishes execution, close f file handle and delete the file file.txt
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
		if err := os.Remove("file.txt"); err != nil {
			log.Fatalln(err)
		}
	}()

	// Get the file's current permissions mode
	info, err := f.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	// Print the file's current permissions mode
	fmt.Println(info.Mode())

	// Now use Chmod to change the files permissions
	//
	// Chmod changes the mode of the file to mode. If there is an error,
	// it will be of type *PathError.
	if err := f.Chmod(os.ModePerm); err != nil {
		log.Fatalln(err)
	}

	// Get the file's current permissions mode again after Chmod
	info, err = f.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	// Print the file's current permissions mode again after Chmod
	fmt.Println(info.Mode())
}

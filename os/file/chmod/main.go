package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a new file with the permissions 0644
	f, _ := os.OpenFile("file.txt", os.O_CREATE|os.O_RDWR, 0644)

	// When main finishes execution, close f file handle and delete the file file.txt
	defer func() {
		f.Close()
		os.Remove("file.txt")
	}()

	// Get the file's current permissions mode
	info, _ := f.Stat()

	// Print the file's current permissions mode
	fmt.Println(info.Mode())

	// Now use Chmod to change the files permissions
	//
	// Chmod changes the mode of the file to mode. If there is an error,
	// it will be of type *PathError.
	f.Chmod(0755)

	// Get the file's current permissions mode again after Chmod
	info, _ = f.Stat()

	// Print the file's current permissions mode again after Chmod
	fmt.Println(info.Mode())
}

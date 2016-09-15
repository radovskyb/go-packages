package main

import "os"

func main() {
	// Create a new file with the permissions 0644
	os.Create("file.txt")

	// Change the file file.txt's permissions to 0755
	//
	// Chmod changes the mode of the named file to mode. If the file
	// is a symbolic link, it changes the mode of the link's target.
	// If there is an error, it will be of type *PathError.
	os.Chmod("file.txt", 0755)
}

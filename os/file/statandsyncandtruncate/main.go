package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the file file.txt for reading and writing
	file, err := os.OpenFile("file.txt", os.O_RDWR, 0755)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Defer to close the file handle file
	defer file.Close()

	// Write some bytes to the file
	_, err = file.Write([]byte("Hello, World! Wassup?"))
	if err != nil {
		log.Fatalln(err)
	}

	// Get the file info description of file.txt
	//
	// Stat returns the FileInfo structure describing file. If there is
	// an error, it will be of type *PathError.
	fileInfo, err := file.Stat()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print the file's size retrieved from file.Stat() before
	// truncating the file below
	fmt.Println(fileInfo.Size())

	// Truncate the file's size to 10 bytes
	//
	// Truncate changes the size of the file. It does not change the
	// I/O offset. If there is an error, it will be of type *PathError.
	os.Truncate(file.Name(), 10)

	// Commit the current contents of the file to stable storage, we would
	// do this for example if the file is really big or incase of a power
	// failure occurring.
	//
	// Sync commits the current contents of the file to stable storage.
	// Typically, this means flushing the file system's in-memory copy
	// of recently written data to disk.
	if err := file.Sync(); err != nil {
		log.Fatalln(err)
	}

	// Now that the file has been truncated, get the file info again
	newInfo, err := file.Stat()

	// Print the file's size retrieved from file.Stat() after truncating it
	fmt.Println(newInfo.Size())
}

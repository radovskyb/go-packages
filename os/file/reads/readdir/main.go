package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the current directory for reading
	f, err := os.Open(".")

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Get the file's directories contents
	//
	// Readdir reads the contents of the directory associated with file and
	// returns a slice of up to n FileInfo values, as would be returned by Lstat,
	// in directory order. Subsequent calls on the same file will yield
	// further FileInfos.
	//
	// If n > 0, Readdir returns at most n FileInfo structures.
	// In this case, if Readdir returns an empty slice, it will return
	// a non-nil error explaining why. At the end of a directory, the error is io.EOF.
	//
	// If n <= 0, Readdir returns all the FileInfo from the directory in a
	// single slice. In this case, if Readdir succeeds (reads all the way
	// to the end of the directory), it returns the slice and a nil error.
	// If it encounters an error before the end of the directory, Readdir
	// returns the FileInfo read until that point and a non-nil error.
	fi, err := f.Readdir(0)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the file info names
	for _, f := range fi {
		fmt.Println(f.Name())
	}

	// We could also call Readdirnames to get the file info names directly.
	//
	// Readdirnames reads and returns a slice of names from the directory f.
	// If n > 0, Readdirnames returns at most n names. In this case,
	// if Readdirnames returns an empty slice, it will return a
	// non-nil error explaining why. At the end of a directory, the error is io.EOF.
	//
	// If n <= 0, Readdirnames returns all the names from the directory in
	// a single slice. In this case, if Readdirnames succeeds (reads all the
	// way to the end of the directory), it returns the slice and a nil error.
	// If it encounters an error before the end of the directory, Readdirnames
	// returns the names read until that point and a non-nil error.
}

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Get the current working directory
	curwd, err := os.Getwd()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Run filepath.Walk on the current working directory with the WalkFunc
	// function walkFn
	//
	// Walk walks the file tree rooted at root, calling walkFn for each file or
	// directory in the tree, including root. All errors that arise visiting files
	// and directories are filtered by walkFn. The files are walked in lexical
	// order, which makes the output deterministic but means that for very
	// large directories Walk can be inefficient.
	// Walk does not follow symbolic links.
	if err := filepath.Walk(curwd, walkFn); err != nil {
		log.Fatalln(err)
	}
}

// Create a function that implements the WalkFunc interface type that for each
// directory that it walks along to, prints out each directory and file name and
// how many bytes in size it is
//
// WalkFunc is the type of the function called for each file or directory
// visited by Walk. The path argument contains the argument to Walk as a
// prefix; that is, if Walk is called with "dir", which is a directory containing
// the file "a", the walk function will be called with argument "dir/a". The info
// argument is the os.FileInfo for the named path.
//
// If there was a problem walking to the file or directory named by path,
// the incoming error will describe the problem and the function can decide
// how to handle that error (and Walk will not descend into that directory).
// If an error is returned, processing stops. The sole exception is when
// the function returns the special value SkipDir. If the function returns
// SkipDir when invoked on a directory, Walk skips the directory's contents
// entirely. If the function returns SkipDir when invoked on a non-directory
// file, Walk skips the remaining files in the containing directory.
func walkFn(path string, fileInfo os.FileInfo, _ error) error {
	fmt.Printf("%s with %d bytes\n", path, fileInfo.Size())
	return nil
}

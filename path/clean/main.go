package main

import (
	"fmt"
	"path"
)

func main() {
	// Create a new string which is the absolute path to this file that also
	// contains some ../'s and a ./ which goes back one directory and
	// then back into it again  etc.. So in  this case path.Clean function
	// will return a cleaned version of fullPath with no ../'s or ./'s
	fullPath := "/Users/Benjamin/Workspace/go/src/../src/pkgchallenge/path/../path/./base"

	// Print the shortest path name equivalent to fullPath above
	//
	// Clean returns the shortest path name equivalent to path by
	// purely lexical processing. It applies the following rules iteratively
	// until no further processing can be done:
	//
	// 1. Replace multiple slashes with a single slash.
	// 2. Eliminate each . path name element (the current directory).
	// 3. Eliminate each inner .. path name element (the parent directory)
	// along with the non-.. element that precedes it.
	// 4. Eliminate .. elements that begin a rooted path:
	// that is, replace "/.." by "/" at the beginning of a path.
	//
	// The returned path ends in a slash only if it is the root "/".
	//
	// If the result of this process is an empty string, Clean returns the string ".".
	//
	// Prints: /Users/Benjamin/Workspace/go/src/pkgchallenge/path/base
	fmt.Println(path.Clean(fullPath))
}

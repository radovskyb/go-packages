package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create 2 new *Files', one from Stdin (0) and one from Stdout (1).
	//
	// NewFile returns a new File with the given file descriptor and name.
	stdin := os.NewFile(0, "Standard Input")
	stdout := os.NewFile(1, "Standard Output")

	// Print out the 2 *File's names receiving their file stats from Stat()
	stdinInfo, err := stdin.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	stdoutInfo, err := stdout.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	// Print stdin's name
	fmt.Println(stdinInfo.Name()) // Standard Input

	// Print stdout's name
	fmt.Println(stdoutInfo.Name()) // Standard Output

	// Now copy from stdin to stdout using io.Copy (Simple echo program)
	_, err = io.Copy(stdout, stdin)
	if err != nil {
		log.Fatalln(err)
	}
}

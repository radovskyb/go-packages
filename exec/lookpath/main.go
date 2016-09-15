package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Use exec.LookPath to search for the full path of the binary
	// of the system command `ls` (/bin/ls)
	//
	// LookPath searches for an executable binary named file
	// in the directories named by the PATH environment variable.
	// If file contains a slash, it is tried directly and the PATH
	// is not consulted.
	//
	// The result may be an absolute path or a path relative to the
	// current directory.
	if executablePath, err := exec.LookPath("ls"); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(executablePath) // /bin/ls
	}
}

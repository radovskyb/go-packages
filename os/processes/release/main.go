package main

import (
	"log"
	"os"
)

func main() {
	// Start a new process and set its stdout and stderr attributes
	// to os.Stdout and os.Stderr but leave it's stdin to nil since no
	// input will be sent to it
	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{nil, os.Stdout, os.Stderr}
	proc, err := os.StartProcess("/bin/ls", []string{"-la"}, &procAttr)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Release any resources associated with the process proc
	//
	// Release releases any resources associated with the Process p,
	// rendering it unusable in the future. Release only needs to
	// be called if Wait is not.
	err = proc.Release()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}
}

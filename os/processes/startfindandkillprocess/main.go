package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Start a new process and set its stdout and stderr attributes
	// to os.Stdout and os.Stderr but leave it's stdin to nil since no
	// input will be sent to it
	//
	// StartProcess starts a new process with the program, arguments and
	// attributes specified by name, argv and attr.
	//
	// StartProcess is a low-level interface. The os/exec package provides
	// higher-level interfaces.
	//
	// If there is an error, it will be of type *PathError.
	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{nil, os.Stdout, os.Stderr}
	proc, err := os.StartProcess("/bin/ls", []string{"-la"}, &procAttr)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Find a process using os.FindProcess
	//
	// FindProcess looks for a running process by its pid. The Process
	// it returns can be used to obtain information about the underlying
	// operating system process.
	process, err := os.FindProcess(proc.Pid)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print the processes pid returned from os.FindProcess
	fmt.Println(process.Pid)

	// Now kill the process straight away
	//
	// Kill causes the Process to exit immediately.
	err = proc.Kill()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}
}

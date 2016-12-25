package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	// Create a new command
	cmd := exec.Command("ls", "-la")

	// Return a pipe that will be connected to the commands standard error
	//
	// StderrPipe returns a pipe that will be connected to the command's
	// standard error when the command starts.
	//
	// Wait will close the pipe after seeing the command exit, so most
	// callers need not close the pipe themselves; however, an implication
	// is that it is incorrect to call Wait before all reads from the pipe
	// have completed. For the same reason, it is incorrect to use Run
	// when using StderrPipe. See the StdoutPipe example for idiomatic usage.
	errpipe, err := cmd.StderrPipe()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Start the command and check for errors and log them if any occur
	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}

	// Create a bytes buffer to store any possible stderr output from cmd
	var buf bytes.Buffer

	// Copy from errpipe into buf
	_, err = io.Copy(&buf, errpipe)
	if err != nil {
		log.Fatalln(err)
	}

	// Run wait on cmd now errpipe has finished being read from
	if err := cmd.Wait(); err != nil {
		log.Fatalln(err)
	}

	// Check if buf is empy and if so print a message that it was empty
	// otherwise print out buf as a string. In this case since the command
	// run would not have returned anything on its standard error pipe,
	// buf will be empty
	if buf.Len() == 0 {
		fmt.Println("No standard error's occurred from running the command.")
	} else {
		fmt.Print(buf.String())
	}
}

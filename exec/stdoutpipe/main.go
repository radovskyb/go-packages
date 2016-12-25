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

	// Return a pipe from the commands standard output before
	// the command is started
	//
	// StdoutPipe returns a pipe that will be connected to the command's
	// standard output when the command starts.
	//
	// Wait will close the pipe after seeing the command exit, so most callers
	// need not close the pipe themselves; however, an implication is that
	// it is incorrect to call Wait before all reads from the pipe have completed.
	// For the same reason, it is incorrect to call Run when using StdoutPipe.
	// See the example for idiomatic usage.
	stdout, err := cmd.StdoutPipe()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}

	// Create a new bytes buffer to store the output of the command
	var buf bytes.Buffer

	// Copy the contents from stdout into bytes buffer buf
	_, err = io.Copy(&buf, stdout)
	if err != nil {
		log.Fatalln(err)
	}

	// Now that everything needed has been read from stdout, call wait
	// on the command and check if any errors occurred whilst calling it
	if err := cmd.Wait(); err != nil {
		// If any errors did occur, log them
		log.Fatalln(err)
	}

	// Print out buf's contents as a string using buf's String() function
	fmt.Println(buf.String())
}

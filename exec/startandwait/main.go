package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Create a new command
	cmd := exec.Command("ls")

	// Make the command's stdout point to os.Stdout so the
	// results of the command are printed to os.Stdout
	cmd.Stdout = os.Stdout

	// Now start the command
	err := cmd.Start()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Now run wait on the command to wait for the command to finish and
	// to release any resources associated with the command
	//
	// Wait waits for the command to exit. It must have been started by Start.
	//
	// The returned error is nil if the command runs, has no problems
	// copying stdin, stdout, and stderr, and exits with a zero exit status.
	//
	// If the command fails to run or doesn't complete successfully,
	// the error is of type *ExitError. Other error types may
	// be returned for I/O problems.
	//
	// If c.Stdin is not an *os.File, Wait also waits for the
	// I/O loop copying from c.Stdin into the process's
	// standard input to complete.
	//
	// Wait releases any resources associated with the Cmd.
	if err := cmd.Wait(); err != nil {
		log.Fatalln(err)
	}
}

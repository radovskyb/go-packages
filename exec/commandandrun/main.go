package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Create a new command
	//
	// Command returns the Cmd struct to execute the named program
	// with the given arguments.
	//
	// It sets only the Path and Args in the returned structure.
	//
	// If name contains no path separators, Command uses LookPath to
	// resolve the path to a complete name if possible. Otherwise it
	// uses name directly.
	//
	// The returned Cmd's Args field is constructed from the command
	// name followed by the elements of arg, so arg should not include
	// the command name itself. For example, Command("echo", "hello")
	cmd := exec.Command("ls", "-la")

	// Set the commands standard output to os.Stdout so the results of
	// the command are printed to os.Stdout
	cmd.Stdout = os.Stdout

	// Run the command and check if there were any errors and
	// if there were, log them
	//
	// Run starts the specified command and waits for it to complete.
	//
	// The returned error is nil if the command runs, has no problems
	// copying stdin, stdout, and stderr, and exits with a zero exit status.
	//
	// If the command fails to run or doesn't complete successfully,
	// the error is of type *ExitError. Other error types may be
	// returned for I/O problems.
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

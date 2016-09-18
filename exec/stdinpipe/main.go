package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Create a new command that will run the ruby file sayhello.rb
	cmd := exec.Command("ruby", "sayhello.rb")

	// Create a new standard input pipe for cmd
	stdinpipe, err := cmd.StdinPipe()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Write the byte slice `Benji` in to the standard input pipe
	_, err = stdinpipe.Write([]byte("Benji"))
	if err != nil {
		log.Fatalln(err)
	}

	// Now close the pipe so that the input will be received on the
	// other end of the pipe because the command being run will not
	// exit until standard input is closed
	//
	// StdinPipe returns a pipe that will be connected to the command's
	// standard input when the command starts. The pipe will be closed
	// automatically after Wait sees the command exit. A caller need only
	// call Close to force the pipe to close sooner. For example, if the
	// command being run will not exit until standard input is closed,
	// the caller must close the pipe.
	if err := stdinpipe.Close(); err != nil {
		log.Fatalln(err)
	}

	// Now run cmd.CombinedOutput() to return both the standard output
	// and standard error into the byte slice output and if there are
	// any errors then log them, otherwise print out the output as a string
	if output, err := cmd.CombinedOutput(); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Print(string(output))
	}
}

package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Store the combined standard output and standard error of the
	// command ls in the byte slice combinedOutput
	//
	// CombinedOutput runs the command and returns its combined
	// standard output and standard error.
	combinedOutput, err := exec.Command("ls").CombinedOutput()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the combined output of the command
	fmt.Print(string(combinedOutput))
}

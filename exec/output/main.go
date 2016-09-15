package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Run the ls command and return it's output
	//
	// Output runs the command and returns its standard output.
	out, err := exec.Command("ls").Output()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the commands output as a string
	fmt.Print(string(out))
}

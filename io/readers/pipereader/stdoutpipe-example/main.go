package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Create a new command to run `ls -la`
	cmd := exec.Command("ls", "-la")

	// Now use cmd.StdoutPipe() to pipe the commands stdout
	// into the variable stdout which underneath will use
	// io.Pipe and cmd.Wait etc to achieve what was achieved
	// with command piping using io.Pipe in the io package
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}

	// Create a new buffered reader `in` that will read from
	// the commands piped stdout
	in := bufio.NewReader(stdout)

	// Now create the second command that will read it's stdin
	// from the first commands stdout, the second command is
	// `grep .go` which will grep through the results of the first
	// commands output and return any results that have `.go` in
	// its line
	cmd = exec.Command("grep", ".go")

	// Use `in` as the commands stdin
	cmd.Stdin = in

	// Redirect the commands stdout to os.Stdout
	// so we can see the results of the second command
	cmd.Stdout = os.Stdout

	// Run the second command
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

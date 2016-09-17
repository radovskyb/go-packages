package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	// Get the process id of the caller
	pid := os.Getpid()

	// Get the process id of the caller's parent
	ppid := os.Getppid()

	// Find the process based on the pid of the caller
	proc1, err := os.FindProcess(pid)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Find the process based on the pid of the caller's parent
	proc2, err := os.FindProcess(ppid)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print the first processes pid returned from find process
	fmt.Println(proc1.Pid)

	// Print the second processes pid returned from find process
	fmt.Println(proc2.Pid)

	// Now send a kill signal to the second process which
	// is the first processes parent, therefore it will end both
	// programs using the kill signal.
	//
	// In this case, it is actually killing the go run main.go process
	if err := proc2.Signal(syscall.SIGKILL); err != nil {
		log.Fatalln(err)
	}
}

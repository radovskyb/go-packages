package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	// Open/create `file.txt`.
	fd, err := syscall.Open("file.txt", syscall.O_WRONLY|syscall.O_CREAT|syscall.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := syscall.Close(fd); err != nil {
			log.Fatalln(err)
		}
	}()

	// Store a copy of the current `stdout`'s `fd`.
	oldStdout, err := syscall.Dup(syscall.Stdout)
	if err != nil {
		log.Fatalln(err)
	}

	// Make `stdout` be a copy of `fd`, therefore writes to `stdout` will
	// actually write to `file.txt`.
	if err := syscall.Dup2(fd, syscall.Stdout); err != nil {
		log.Fatalln(err)
	}

	// Print some text to `stdout` which will actually be printing to `file.txt`
	fmt.Println("Hello, World! 1")
	fmt.Println("Hello, World! 2")
	fmt.Println("Hello, World! 3")

	// Now make `fd` 1 (current `stdout` pointing to `fd`) a copy of the old `stdout`.
	if err := syscall.Dup2(oldStdout, syscall.Stdout); err != nil {
		log.Fatalln(err)
	}

	// Close the saved `stdout`'s `fd` now that `stdout` is back to normal.
	if err := syscall.Close(oldStdout); err != nil {
		log.Fatalln(err)
	}

	// Print some text to `stdout`
	fmt.Println("Back to normal.")
}

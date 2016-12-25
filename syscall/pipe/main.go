package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	// Create an []int to store 2 fd's. One for the pipe
	// reader's fd and one for the pipe writer's fd.
	fds := make([]int, 2)

	// Create a pipe reader and writer and store their fd's in `fds`.
	if err := syscall.Pipe(fds); err != nil {
		log.Fatalln(err)
	}

	go func() {
		// Write some bytes into the pipe writer.
		_, err := syscall.Write(fds[1], []byte("Hello!"))
		if err != nil {
			log.Fatalln(err)
		}

		// Close the pipe writer, fds[1].
		if err := syscall.Close(fds[1]); err != nil {
			log.Fatalln(err)
		}
	}()

	// Create a []byte to buffer 2 bytes at a time from
	// the pipe. (can make buffer value higher)
	buf := make([]byte, 2)
	for {
		// Read from the pipe into `buf`.
		n, err := syscall.Read(fds[0], buf)

		// If n is zero or an error occurred, exit the loop, we are
		// finished reading from the pipe.
		if n == 0 || err != nil {
			break
		}

		// Print the bytes read from the pipe,
		// in string format, to stdout.
		fmt.Printf("%s", buf[0:n])
	}

	// Close the pipe reader, fds[0].
	if err := syscall.Close(fds[0]); err != nil {
		log.Fatalln(err)
	}

	fmt.Println()
}

package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Create a first command to run `ls`
	first := exec.Command("ls", "-la")

	// Create a second command that we want to run
	// after piping the data from the first command into it.
	// The command is `grep main`, which will first need
	// input to run
	second := exec.Command("grep", "main")

	// Create a new pipe reader and pipe writer using io.Pipe()
	//
	// A PipeReader is the read half of a pipe.
	// A PipeWriter is the write half of a pipe.
	//
	// Pipe creates a synchronous in-memory pipe. It can be used to
	// connect code expecting an io.Reader with code expecting an
	// io.Writer. Reads on one end are matched with writes on the other,
	// copying data directly between the two; there is no internal buffering.
	// It is safe to call Read and Write in parallel with each other or
	// with Close. Close will complete once pending I/O is done. Parallel
	// calls to Read, and parallel calls to Write, are also safe:
	// the individual calls will be gated sequentially.
	pReader, pWriter := io.Pipe()

	// Set the first command's Stdout to the pipe writer
	// which we will use to push the commands output on to
	// the pipe, for the pipe reader to read from
	first.Stdout = pWriter

	// Set the second command's Stdin to the pipe reader, so
	// we can then read from the pipe which will have content
	// written into it from the first commands output, written
	// there by the pipe writer
	second.Stdin = pReader

	// Set the second command's stdout to os.Stdout so we can
	// have the second command's printed directly to os.Stdout
	// as it runs
	second.Stdout = os.Stdout

	// Start both commands
	if err := first.Start(); err != nil {
		log.Fatalln(err)
	}
	if err := second.Start(); err != nil {
		log.Fatalln(err)
	}

	// Run Wait on the first command because here we need
	// to wait for the commands input to finish copying
	// to the processes stdin
	//
	// Wait waits for the command to exit.
	// It must have been started by Start.
	//
	// If c.Stdin is not an *os.File, Wait also waits for the I/O loop
	// copying from c.Stdin into the process's standard input
	// to complete.
	//
	// Wait releases any resources associated with the Cmd.
	if err := first.Wait(); err != nil {
		log.Fatalln(err)
	}

	// Close the pipe writer so the pipe reader can grab
	// the contents from it
	if err := pWriter.Close(); err != nil {
		log.Fatalln(err)
	}

	// Run Wait on the first command to free up any
	// resources associated with the second command
	if err := second.Wait(); err != nil {
		log.Fatalln(err)
	}

	// Close the pipe reader now that it's job is done
	if err := pReader.Close(); err != nil {
		log.Fatalln(err)
	}
}

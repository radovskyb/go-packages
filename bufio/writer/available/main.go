package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a new buffered writer to os.Stdout using bufio.NewWriter
	w := bufio.NewWriter(os.Stdout)

	// Print out how many bytes are available in the w buffer
	//
	// Available returns how many bytes are unused in the buffer.
	fmt.Println(w.Available())
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello\n"

	// Create a new buffered reader br that reads from
	// a new strings reader that reads from string s
	br := bufio.NewReader(strings.NewReader(s))

	// Write the buffered reader br's contents to os.Stdout
	//
	// WriteTo implements io.WriterTo.
	n, err := br.WriteTo(os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Wrote %d bytes\n", n)
}

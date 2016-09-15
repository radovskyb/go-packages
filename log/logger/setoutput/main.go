package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Logged: ", log.Lshortfile)

	logger.Print("Hello, Logger (os.Stdout)!")

	var buf bytes.Buffer

	// Change the logger output to log to &buf instead of os.Stdout
	//
	// SetOutput sets the output destination for the logger.
	logger.SetOutput(&buf)

	logger.Print("Hello, Logger (&buf)!")

	fmt.Print(&buf)
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Save os.Stdout for renormalizing os.Stdout after using os.Pipe
	savedStdout := os.Stdout

	// Create a new pipe reader and writer from os.Pipe
	r, w, _ := os.Pipe()

	// Send os.Stdout to the pipe writer
	os.Stdout = w

	// Print some data to os.Stdout which is this case will send to w
	fmt.Printf("Hello from data passed to stdout which is currently w")

	// Close the writer so we can capture it's output and stop it's blocking
	w.Close()

	// Store all the data received from the writer on the reader
	data, _ := ioutil.ReadAll(r)

	// Go back to normal os.Stdout
	os.Stdout = savedStdout

	// Could also have used this line without assigning a
	// savedStdout variable above:  os.Stdout = os.NewFile(1, "/dev/stdout")

	// Print the data received from the reader as a string
	fmt.Println(string(data))
}

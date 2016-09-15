package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World!"

	// Create a new strings reader sr that reads from string s
	sr := strings.NewReader(s)

	// Read all the contents from string reader sr
	//
	// ReadAll reads from r until an error or EOF and returns the
	// data it read. A successful call returns err == nil, not err == EOF.
	// Because ReadAll is defined to read from src until EOF, it does
	// not treat an EOF from Read as an error to be reported.
	read, err := ioutil.ReadAll(sr)
	if err != nil {
		log.Fatalln(err)
	}

	// Print out read variable
	fmt.Println(string(read))
}

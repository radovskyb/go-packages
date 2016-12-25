package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const filename = "file.txt"

func main() {
	// Open the file `file.txt`
	file, err := os.Open(filename)

	// Check if there were any errors opening
	// the file and if there was any, log them
	if err != nil {
		log.Fatal(err)
	}

	// Make sure we close the file handle
	// when main's execution ends
	defer file.Close()

	// Create a new buffered scanner for the file
	fs := bufio.NewScanner(file)

	// Scan through each line in the file
	for fs.Scan() {
		// Print out each line scanned in
		fmt.Println(fs.Text())
	}

	// Check if any scanner errors occurred and if so log them
	if err := fs.Err(); err != nil {
		log.Fatalln(err)
	}
}

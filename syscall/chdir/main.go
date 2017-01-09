package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	// Print the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cwd)

	// Change to the current directory's parent directory.
	if err := syscall.Chdir(".."); err != nil {
		log.Fatalln(err)
	}

	// Print the current working directory again.
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cwd)
}

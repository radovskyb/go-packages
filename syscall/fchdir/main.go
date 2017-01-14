package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"syscall"
)

func main() {
	fd, err := syscall.Open("tmp", syscall.O_RDONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer syscall.Close(fd)

	// Change the current working directory to tmp from it's fd.
	//
	// fchdir is identical to chdir except that the only difference is
	// that the directory is given as an open file descriptor and not by
	// it's path name.
	if err := syscall.Fchdir(fd); err != nil {
		log.Fatalln(err)
	}

	// Read file.txt from the now new current working directory.
	slurp, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s", slurp)
}

package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	// Open main.go (the current file).
	fd, err := syscall.Open("main.go", syscall.O_RDONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	stat := new(syscall.Stat_t)

	// fstat is identical to stat, except that the file to be
	// stat-ed is specified by the file descriptor fd.
	err = syscall.Fstat(fd, stat)
	if err != nil {
		log.Fatalln(err)
	}

	// Print main.go's file size and uid.
	fmt.Println(stat.Size)
	fmt.Println(stat.Uid)
}

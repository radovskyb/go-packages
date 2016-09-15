package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	statfs := new(syscall.Statfs_t)
	if err := syscall.Statfs("/Users/Benjamin", statfs); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(statfs.Bfree)
	fmt.Println(statfs.Owner)
}

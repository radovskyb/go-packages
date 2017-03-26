package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	// Getgroups returns the group IDs of the calling
	// process.
	gids, err := syscall.Getgroups()
	if err != nil {
		log.Fatalln(err)
	}

	// Print out each of the group id's.
	for _, gid := range gids {
		fmt.Println(gid)
	}
}

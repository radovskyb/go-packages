package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	// Get the current user on the computer
	//
	// Current returns the current user.
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	// Print the users name
	fmt.Println(u.Name)
}

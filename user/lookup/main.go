package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	// Lookup a user on the computer by username
	//
	// Lookup looks up a user by username. If the user cannot be found,
	// the returned error is of type UnknownUserError.
	u, err := user.Lookup("Benjamin")
	if err != nil {
		log.Fatalln(err)
	}

	// Print the users home directory
	fmt.Println(u.HomeDir)
}

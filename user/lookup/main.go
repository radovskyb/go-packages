package main

import (
	"fmt"
	"os/user"
)

func main() {
	// Lookup a user on the computer by username
	//
	// Lookup looks up a user by username. If the user cannot be found,
	// the returned error is of type UnknownUserError.
	u, _ := user.Lookup("Benjamin")

	// Print the users home directory
	fmt.Println(u.HomeDir)
}

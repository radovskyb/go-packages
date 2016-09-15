package main

import (
	"fmt"
	"os/user"
)

func main() {
	// Get the current user on the computer
	//
	// Current returns the current user.
	u, _ := user.Current()

	// Print the users name
	fmt.Println(u.Name)
}

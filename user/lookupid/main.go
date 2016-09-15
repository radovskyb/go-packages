package main

import (
	"fmt"
	"os/user"
)

func main() {
	// Lookup a user on the computer by id
	//
	// LookupId looks up a user by userid. If the user cannot be
	// found, the returned error is of type UnknownUserIdError.
	u, _ := user.LookupId("501")

	// Print the users gid
	fmt.Println(u.Gid)
}

package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	// Lookup a user on the computer by id
	//
	// LookupId looks up a user by userid. If the user cannot be
	// found, the returned error is of type UnknownUserIdError.
	u, err := user.LookupId("501")
	if err != nil {
		log.Fatalln(err)
	}

	// Print the users gid
	fmt.Println(u.Gid)
}

package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func main() {
	// Create a new file with the permissions 0644
	f, err := os.OpenFile("file.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	// When main finishes execution, close f file handle and delete the file file.txt
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
		if err := os.Remove("file.txt"); err != nil {
			log.Fatalln(err)
		}
	}()

	// Get the file's current permissions mode
	info, err := f.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	// Print the file's current user's id
	fmt.Println(info.Sys().(*syscall.Stat_t).Uid)

	// Print the file's current group id
	fmt.Println(info.Sys().(*syscall.Stat_t).Gid)

	// Get the current calling user's gid and uid
	user, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	// Convert uid and gid to integers from strings to work with Chown function
	uid, err := strconv.Atoi(user.Uid)
	if err != nil {
		log.Fatalln(err)
	}
	gid, err := strconv.Atoi(user.Gid)
	if err != nil {
		log.Fatalln(err)
	}

	// Change the files uid and gid to the current calling users uid and gid
	if err := f.Chown(uid, gid); err != nil {
		log.Fatalln(err)
	}

	// Get the file's current permissions mode
	info, err = f.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	// Print the file's current user's id
	fmt.Println(info.Sys().(*syscall.Stat_t).Uid)

	// Print the file's current group id
	fmt.Println(info.Sys().(*syscall.Stat_t).Gid)
}

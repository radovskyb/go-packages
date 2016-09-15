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
	f, _ := os.OpenFile("file.txt", os.O_CREATE|os.O_RDWR, 0644)

	// When main finishes execution, close f file handle and delete the file file.txt
	defer func() {
		f.Close()
		os.Remove("file.txt")
	}()

	// Get the file's current permissions mode
	info, _ := f.Stat()

	// Print the file's current user's id
	fmt.Println(info.Sys().(*syscall.Stat_t).Uid)

	// Print the file's current group id
	fmt.Println(info.Sys().(*syscall.Stat_t).Gid)

	// Get the current calling user's gid and uid
	user, _ := user.Current()

	// Convert uid and gid to integers from strings to work with Chown function
	uid, _ := strconv.Atoi(user.Uid)
	gid, _ := strconv.Atoi(user.Gid)

	// Change the files uid and gid to the current calling users uid and gid
	err := f.Chown(uid, gid)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Get the file's current permissions mode
	info, _ = f.Stat()

	// Print the file's current user's id
	fmt.Println(info.Sys().(*syscall.Stat_t).Uid)

	// Print the file's current group id
	fmt.Println(info.Sys().(*syscall.Stat_t).Gid)
}

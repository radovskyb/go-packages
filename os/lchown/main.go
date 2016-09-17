package main

import (
	"log"
	"os"
)

func main() {
	// Change the numeric uid and gid of the file file.txt to the current
	// users uid and gid respectively
	//
	// Lchown changes the numeric uid and gid of the named file. If the
	// file is a symbolic link, it changes the uid and gid of the link
	// itself. If there is an error, it will be of type *PathError.
	if err := os.Lchown("file.txt", os.Getuid(), os.Getgid()); err != nil {
		log.Fatalln(err)
	}
}

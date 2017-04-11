package main

import (
	"fmt"
	"log"
	"syscall"
	"time"
)

func main() {
	stat := new(syscall.Stat_t)

	// Lstat is identical to Stat, except that if path is a symbolic link,
	// then the link itself is stat-ed, not the file that it refers to.
	if err := syscall.Stat("main.go", stat); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(time.Unix(stat.Atimespec.Unix()))
	fmt.Println(time.Unix(stat.Birthtimespec.Unix()))
	fmt.Println(time.Unix(stat.Ctimespec.Unix()))
	fmt.Println(time.Unix(stat.Mtimespec.Unix()))
	fmt.Println(stat.Size)
	fmt.Println(stat.Uid)
}

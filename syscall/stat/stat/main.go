package main

import (
	"fmt"
	"syscall"
	"time"
)

func main() {
	stat := new(syscall.Stat_t)

	syscall.Stat("main.go", stat)

	fmt.Println(time.Unix(stat.Atimespec.Unix()))
	fmt.Println(time.Unix(stat.Birthtimespec.Unix()))
	fmt.Println(time.Unix(stat.Ctimespec.Unix()))
	fmt.Println(time.Unix(stat.Mtimespec.Unix()))
	fmt.Println(stat.Size)
	fmt.Println(stat.Uid)
}

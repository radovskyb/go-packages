package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	fmt.Println(syscall.Getpid())

	pid, _, err := syscall.RawSyscall(syscall.SYS_GETPID, 0, 0, 0)
	if err != 0 {
		log.Fatalln(err)
	}
	fmt.Println(pid)
}

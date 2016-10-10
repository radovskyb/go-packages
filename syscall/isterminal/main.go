package main

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

const path = "/dev/ttys000"

func main() {
	fd, err := syscall.Open(path, syscall.O_RDONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := syscall.Close(fd); err != nil {
			log.Fatalln(err)
		}
	}()

	term := new(syscall.Termios)
	_, _, e0 := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd),
		syscall.TIOCGETA, uintptr(unsafe.Pointer(term)), 0, 0, 0)
	if e0 != 0 {
		log.Fatalln(e0)
	}

	fmt.Printf("%s is a terminal\n", path)
}

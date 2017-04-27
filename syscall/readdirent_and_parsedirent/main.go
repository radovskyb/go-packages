package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	// Open the directory test_dir.
	fd, err := syscall.Open("test_dir", syscall.O_RDONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer syscall.Close(fd)

	// Read directory entries into buf.
	buf := make([]byte, 1024)
	_, err = syscall.ReadDirent(fd, buf)
	if err != nil {
		log.Fatalln(err)
	}

	// Parse the directory entries from buf and return the amount
	// of characters consumed, the count of entries and the names
	// of each entry.
	consumed, count, names := syscall.ParseDirent(buf, -1, []string{})
	fmt.Printf("Consumed: %d bytes\nEntries: %d\n", consumed, count)

	fmt.Println("\nNames:")
	for _, name := range names {
		fmt.Println(name)
	}
}

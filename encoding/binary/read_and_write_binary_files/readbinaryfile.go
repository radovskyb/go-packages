package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("binaryfile")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	var pi float64
	err = binary.Read(f, binary.LittleEndian, &pi)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(pi)
}

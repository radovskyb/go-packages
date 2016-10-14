package main

import (
	"encoding/binary"
	"log"
	"os"
)

func main() {
	f, err := os.Create("binaryfile")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	err = binary.Write(f, binary.LittleEndian, b)
	if err != nil {
		log.Fatalln(err)
	}
}

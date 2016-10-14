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

	fl := 123.456
	err = binary.Write(f, binary.LittleEndian, fl)
	if err != nil {
		log.Fatalln(err)
	}

	fl = 789.101
	err = binary.Write(f, binary.LittleEndian, fl)
	if err != nil {
		log.Fatalln(err)
	}

	fl = 121.314
	err = binary.Write(f, binary.LittleEndian, fl)
	if err != nil {
		log.Fatalln(err)
	}
}

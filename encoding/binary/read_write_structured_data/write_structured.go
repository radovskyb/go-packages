package main

import (
	"encoding/binary"
	"log"
	"os"
)

type Integers struct {
	i1 uint16
	i2 int32
	i3 int64
}

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

	i := Integers{i1: 16, i2: 32, i3: 64}
	err = binary.Write(f, binary.LittleEndian, i)
	if err != nil {
		log.Fatalln(err)
	}
}

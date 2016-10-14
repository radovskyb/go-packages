package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

type Integers struct {
	I1 uint16
	I2 int32
	I3 int64
}

func main() {
	f, err := os.Open("binaryfile")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	i := Integers{}
	err = binary.Read(f, binary.LittleEndian, &i)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(i)
}

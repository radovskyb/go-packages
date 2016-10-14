package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("binaryfile")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	for {
		var val float64
		err = binary.Read(f, binary.LittleEndian, &val)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		fmt.Println(val)
	}
}

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func main() {
	var emptybuf []byte

	buf := []byte{
		144, 192, 192, 129, 132, 136, 140, 144, 16, 0, 1, 1,
	}

	overflowbuf := []byte{
		144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1,
	}

	// ReadUvarint reads an encoded unsigned integer from r and
	// returns it as a uint64.
	num, err := binary.ReadUvarint(bytes.NewBuffer(emptybuf))
	errCheck(err)
	fmt.Println(num)

	num, err = binary.ReadUvarint(bytes.NewBuffer(buf))
	errCheck(err)
	fmt.Println(num)

	num, err = binary.ReadUvarint(bytes.NewBuffer(overflowbuf))
	errCheck(err)
	fmt.Println(num)
}

func errCheck(err error) {
	if err != nil {
		if err != io.EOF {
			fmt.Println(err)
		}
	}
}

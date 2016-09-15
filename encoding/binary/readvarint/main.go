package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func main() {
	var emptybuf []byte

	var buf []byte = []byte{
		144, 192, 192, 129, 132, 136, 140, 144, 16, 0, 1, 1,
	}

	var overflowbuf []byte = []byte{
		144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1,
	}

	// ReadVarint reads an encoded signed integer from r and
	// returns it as an int64.
	num, err := binary.ReadVarint(bytes.NewBuffer(emptybuf))
	errCheck(err)
	fmt.Println(num)

	num, err = binary.ReadVarint(bytes.NewBuffer(buf))
	errCheck(err)
	fmt.Println(num)

	num, err = binary.ReadVarint(bytes.NewBuffer(overflowbuf))
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

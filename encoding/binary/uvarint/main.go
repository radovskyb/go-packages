package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var emptybuf []byte

	buf := []byte{
		144, 192, 192, 129, 132, 136, 140, 144, 16, 0, 1, 1,
	}

	overflowbuf := []byte{
		144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1,
	}

	// Uvarint decodes a uint64 from buf and returns that value and the
	// number of bytes read (> 0). If an error occurred, the value is 0
	// and the number of bytes n is <= 0 meaning:
	//
	//	n == 0: buf too small
	//	n  < 0: value larger than 64 bits (overflow)
	//              and -n is the number of bytes read
	num, ret := binary.Uvarint(emptybuf)
	fmt.Println(num, ret)

	num, ret = binary.Uvarint(buf)
	fmt.Println(num, ret)

	num, ret = binary.Uvarint(overflowbuf)
	fmt.Println(num, ret)
}

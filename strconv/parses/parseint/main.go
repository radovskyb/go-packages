package main

import (
	"fmt"
	"strconv"
)

func main() {
	v32 := "-354634382"

	// Convert the string's integer representation from v32 into
	// a base 10 integer of 32 bits and print it's value and type
	//
	// ParseInt interprets a string s in the given base (2 to 36) and
	// returns the corresponding value i.  If base == 0, the base is
	// implied by the string's prefix: base 16 for "0x", base 8 for
	// "0", and base 10 otherwise.
	//
	// The bitSize argument specifies the integer type
	// that the result must fit into.  Bit sizes 0, 8, 16, 32, and 64
	// correspond to int, int8, int16, int32, and int64.
	//
	// The errors that ParseInt returns have concrete type *NumError
	// and include err.Num = s.  If s is empty or contains invalid
	// digits, err.Err = ErrSyntax and the returned value is 0;
	// if the value corresponding to s cannot be represented by a
	// signed integer of the given size, err.Err = ErrRange and the
	// returned value is the maximum magnitude integer of the
	// appropriate bitSize and sign.
	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	} else {
		fmt.Println(err)
	}

	// Convert the string's integer representation from v32 into
	// a base 16 integer of 32 bits and print it's value and type
	if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	} else {
		fmt.Println(err)
	}

	v64 := "-3546343826724305832"

	// Convert the string's integer representation from v64 into
	// a base 10 integer of 64 bits and print it's value and type
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	} else {
		fmt.Println(err)
	}

	// Convert the string's integer representation from v64 into
	// a base 16 integer of 64 bits and print it's value and type
	if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	} else {
		fmt.Println(err)
	}
}

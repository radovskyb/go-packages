package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := 3.1415926535

	// FormatFloat converts the floating-point number f to a string,
	// according to the format fmt and precision prec.  It rounds the
	// result assuming that the original was obtained from a floating-point
	// value of bitSize bits (32 for float32, 64 for float64).
	//
	// The format fmt is one of
	// 'b' (-ddddp±ddd, a binary exponent),
	// 'e' (-d.dddde±dd, a decimal exponent),
	// 'E' (-d.ddddE±dd, a decimal exponent),
	// 'f' (-ddd.dddd, no exponent),
	// 'g' ('e' for large exponents, 'f' otherwise), or
	// 'G' ('E' for large exponents, 'f' otherwise).
	//
	// The precision prec controls the number of digits
	// (excluding the exponent) printed by the 'e', 'E', 'f', 'g', and 'G' formats.
	// For 'e', 'E', and 'f' it is the number of digits after the decimal point.
	// For 'g' and 'G' it is the total number of digits.
	// The special precision -1 uses the smallest number of digits
	// necessary such that ParseFloat will return f exactly.
	s32 := strconv.FormatFloat(v, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)

	s64 := strconv.FormatFloat(v, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)
}

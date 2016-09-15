package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Create a string that contains a float value
	v := "3.1415926535"

	// Convert and store the flaot value represented in the string v
	// into float variable s
	//
	// ParseFloat converts the string s to a floating-point number
	// with the precision specified by bitSize: 32 for float32, or 64
	// for float64.
	// When bitSize=32, the result still has type float64, but it will be
	// convertible to float32 without changing its value.
	//
	// If s is well-formed and near a valid floating point number,
	// ParseFloat returns the nearest floating point number rounded
	// using IEEE754 unbiased rounding.
	//
	// The errors that ParseFloat returns have concrete type *NumError
	// and include err.Num = s.
	//
	// If s is not syntactically well-formed, ParseFloat returns
	// err.Err = ErrSyntax.
	//
	// If s is syntactically well-formed but is more than 1/2 ULP
	// away from the largest floating point number of the given size,
	// ParseFloat returns f = ±Inf, err.Err = ErrRange.
	if s, err := strconv.ParseFloat(v, 32); err == nil {
		// Print out the float value in s and it's type
		fmt.Printf("%T, %v\n", s, s)
	}

	if s, err := strconv.ParseFloat(v, 64); err == nil {
		// Print out the float value in s and it's type
		fmt.Printf("%T, %v\n", s, s)
	}
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Create a string str that isn't a string representation
	// of an integer
	str := "Not a number"

	// Try to call strconv.ParseFloat on str, however since there
	// is no string representation of an integer inside of str,
	// the call will return an error which will be of type NumError
	//
	// A NumError records a failed conversion.
	if _, err := strconv.ParseFloat(str, 64); err != nil {
		e := err.(*strconv.NumError) // <- NumError
		fmt.Println("Func:", e.Func)
		fmt.Println("Num:", e.Num)
		fmt.Println("Err:", e)
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create a new integer named integer
	var integer int

	// Create a new reflect.Value object from integer
	//
	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i. ValueOf(nil) returns the zero Value.
	iValue := reflect.ValueOf(integer)

	// Print out iValues int value
	fmt.Println(iValue.Int()) // 0
}

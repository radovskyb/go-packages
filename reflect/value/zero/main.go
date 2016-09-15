package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create a reflect.Value object with it's type's value
	// set to it's types zero value, in this case, 0 for an integer
	//
	// Zero returns a Value representing the zero value for
	// the specified type. The result is different from the zero
	// value of the Value struct, which represents no value at all.
	// For example, Zero(TypeOf(42)) returns a Value with Kind Int and
	// value 0. The returned value is neither addressable nor settable.
	zeroedInteger := reflect.Zero(reflect.TypeOf((*int)(nil)).Elem())

	// Print out zeroedInteger
	fmt.Println(zeroedInteger)
}

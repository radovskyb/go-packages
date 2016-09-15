package main

import (
	"fmt"
	"reflect"
)

func main() {
	// swap is the implementation passed to MakeFunc.
	// It must work in terms of reflect.Values so that it is possible
	// to write code without knowing beforehand what the types
	// will be.
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

	// makeSwap expects fptr to be a pointer to a nil function.
	// It sets that pointer to a new function created with MakeFunc.
	// When the function is invoked, reflect turns the arguments
	// into Values, calls swap, and then turns swap's result slice
	// into the values returned by the new function.
	makeSwap := func(fptr interface{}) {
		// fptr is a pointer to a function.
		// Obtain the function value itself (likely nil)
		// as a reflect.Value so that we can query its type
		// and then set the value.
		fn := reflect.ValueOf(fptr).Elem()

		// Make a function of the right type
		//
		// MakeFunc returns a new function of the given Type
		// that wraps the function fn. When called, that new function
		// does the following:
		//
		//	- converts its arguments to a slice of Values.
		//	- runs results := fn(args).
		//	- returns the results as a slice of Values, one
		//	  per formal result.
		//
		// The implementation fn can assume that the argument Value slice
		// has the number and type of arguments given by typ.
		// If typ describes a variadic function, the final Value is itself
		// a slice representing the variadic arguments, as in the
		// body of a variadic function. The result Value slice
		// returned by fn must have the number and type of results
		// given by typ.
		//
		// The Value.Call method allows the caller to invoke a
		// typed function in terms of Values; in contrast, MakeFunc
		// allows the caller to implement a typed function in terms
		// of Values.
		v := reflect.MakeFunc(fn.Type(), swap)

		// Assign it to the value fn represents.
		fn.Set(v)
	}

	// Make and call a swap function for ints.
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))

	// Make and call a swap function for float64s.
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(4.56, 1.23))
}

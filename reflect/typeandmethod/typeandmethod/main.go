package main

import (
	"fmt"
	"reflect"
)

// Create a new integer type `myInt`
type myIntType int

func (m myIntType) MyIntsValue() string {
	return string(m)
}

func main() {
	// Create a new myIntType variable `myInt`
	myInt := myIntType(10)

	// Get the type of `myInt` (returns a reflect.Type object)
	//
	// Type is the representation of a Go type.
	//
	// Not all methods apply to all kinds of types.  Restrictions,
	// if any, are noted in the documentation for each method.
	// Use the Kind method to find out the kind of type before
	// calling kind-specific methods.  Calling a method
	// inappropriate to the kind of type causes a run-time panic.
	typeofMyInt := reflect.TypeOf(myInt)

	// Print out typeofMyInt
	fmt.Println(typeofMyInt) // int

	// reflect.Type methods:

	// Align returns the alignment in bytes of a value of
	// this type when allocated in memory.
	fmt.Println(typeofMyInt.Align())

	// FieldAlign returns the alignment in bytes of a value of
	// this type when used as a field in a struct.
	fmt.Println(typeofMyInt.FieldAlign())

	// Print the number of methods that myInt has in it's method set
	//
	// NumMethod returns the number of methods in the type's method set.
	fmt.Println(typeofMyInt.NumMethod())

	// Print the typeofMyInt's type's name
	//
	// Name returns the type's name within its package.
	// It returns an empty string for unnamed types.
	fmt.Println(typeofMyInt.Name()) // myIntType

	// PkgPath returns a named type's package path, that is, the import path
	// that uniquely identifies the package, such as "encoding/base64".
	// If the type was predeclared (string, error) or unnamed
	// (*T, struct{}, []int), the package path will be the empty string.
	fmt.Println(typeofMyInt.PkgPath())

	// Size returns the number of bytes needed to store
	// a value of the given type; it is analogous to unsafe.Sizeof.
	fmt.Println(typeofMyInt.Size())

	// String returns a string representation of the type.
	// The string representation may use shortened package names
	// (e.g., base64 instead of "encoding/base64") and is not
	// guaranteed to be unique among types.  To test for equality,
	// compare the Types directly.
	fmt.Println(typeofMyInt.String())

	// Print the types kind
	//
	// Kind returns the specific kind of this type.
	fmt.Println(typeofMyInt.Kind())

	// Comparable reports whether values of this type are comparable.
	fmt.Println(typeofMyInt.Comparable())

	// Bits returns the size of the type in bits.
	// It panics if the type's Kind is not one of the
	// sized or unsized Int, Uint, Float, or Complex kinds.
	fmt.Println(typeofMyInt.Bits())

	// Get myIntType's first method (returns a reflect.Method object)
	//
	// Method returns the i'th method in the type's method set.
	// It panics if i is not in the range [0, NumMethod()).
	//
	// For a non-interface type T or *T, the returned Method's Type
	// and Func fields describe a function whose first argument is
	// the receiver.
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver, and the Func field is nil.
	method := typeofMyInt.Method(0)

	// reflect.Method fields:
	fmt.Println()

	// Print the method's function (returns a reflect.Value)
	fmt.Println(&method.Func)

	// Print the method's index
	fmt.Println(method.Index) // 0 (it's the first method for myIntType)

	// Print the name of the method
	fmt.Println(method.Name) // MyIntsValue

	// Print the method's type
	fmt.Println(method.Type) // func(main.myIntType) string

	fmt.Println()

	// Get a myInt's method by name (Same as the above: returns a
	// reflect.Method but also returns whether it was found or not too)
	//
	// MethodByName returns the method with that name in the type's
	// method set and a boolean indicating if the method was found.
	//
	// For a non-interface type T or *T, the returned Method's Type
	// and Func fields describe a function whose first argument
	// is the receiver.
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver, and the Func field is nil.
	mbn, found := typeofMyInt.MethodByName("MyIntsValue")
	if found {
		fmt.Println("Method MyIntsValue exists")
	}
	fmt.Println(mbn)
}

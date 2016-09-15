package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

type myType interface {
	myMethod()
}

type myObject struct {
}

func (mo myObject) myMethod() {
	fmt.Println("myMethod")
}

func main() {
	// Get the type of io.Writer interface
	//
	// As interface types are only used for static typing, a
	// common idiom to find the reflection Type for an interface
	// type Foo is to use a *Foo value.
	//
	// TypeOf returns the reflection Type that represents the dynamic
	// type of i. If i is a nil interface value, TypeOf returns nil.
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()

	// Get the type of os.File interface
	fileType := reflect.TypeOf((*os.File)(nil))

	// Check if the type of an os.File struct object implements
	// the interface io.Writer
	fmt.Println(fileType.Implements(writerType)) // true

	// Get the type of the myType interface
	myTypeType := reflect.TypeOf((*myType)(nil)).Elem()

	// Get the type of a myObject object
	myObjectType := reflect.TypeOf((*myObject)(nil))

	// Print whether the myObject object type implements
	// the interface myType or not
	fmt.Println(myObjectType.Implements(myTypeType)) // true
}

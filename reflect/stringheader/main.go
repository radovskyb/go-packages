package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// Create a new string
	str := "Hello, World!"

	// Create a new reflect.StringHeader variable from `str`, `sh`
	//
	// StringHeader is the runtime representation of a string.
	// It cannot be used safely or portably and its representation may
	// change in a later release.
	// Moreover, the Data field is not sufficient to guarantee the data
	// it references will not be garbage collected, so programs must keep
	// a separate, correctly typed pointer to the underlying data.
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))

	fmt.Println("str's len:", sh.Len) // 13
	fmt.Println("str's data:", sh.Data)
}

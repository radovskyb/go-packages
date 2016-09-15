package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// Create a new []string `slice`
	slice := []string{"abc", "def"}

	// Create a new reflect.SliceHeader variable from `slice`, `sh`
	//
	// SliceHeader is the runtime representation of a slice.
	// It cannot be used safely or portably and its representation may
	// change in a later release.
	// Moreover, the Data field is not sufficient to guarantee the data
	// it references will not be garbage collected, so programs must keep
	// a separate, correctly typed pointer to the underlying data.
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	fmt.Println("slice's length:", sh.Len)   // 2
	fmt.Println("slice's capacity:", sh.Cap) // 2
	fmt.Println("slice's data:", sh.Data)
}

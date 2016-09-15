package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str chan string // channel string type

	var strVal reflect.Value = reflect.ValueOf(str)

	// Create a new channel
	//
	// MakeChan creates a new channel with the specified type
	// and buffer size.
	value := reflect.MakeChan(strVal.Type(), 1)

	fmt.Printf("Value type is [%v] and capacity [%v] byte/s.\n",
		value.Kind(), value.Cap())
}

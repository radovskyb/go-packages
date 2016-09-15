package main

import (
	"fmt"
	"reflect"
)

type t struct {
	Fi int
	fs string
}

var r *t = &t{5, "Hi."}

var i64 int64 = 10

func main() {
	v := reflect.Indirect(reflect.ValueOf(r))

	fmt.Println("r's field names, exported and unexported:")
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Type().Field(i).Name)
	}

	fmt.Println("\nr's field values:")
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf(
			"%v: %v\n",
			v.Field(i).Type(),
			v.Field(i),
		)
	}

	fmt.Println("\nSettable fields:")
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf(
			"%s: %t\n",
			v.Type().Field(i).Name,
			v.Field(i).CanSet(),
		)
	}

	fmt.Println("\nSetting field Fi to a new integer value...")
	v.Field(0).SetInt(i64)

	fmt.Println("\nField Fi's new value:")
	fmt.Println(v.Field(0))
}

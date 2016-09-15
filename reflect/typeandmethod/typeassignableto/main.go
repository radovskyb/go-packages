package main

import (
	"fmt"
	"reflect"
)

type myInt int

func main() {
	a := 10
	t1 := reflect.TypeOf(a)

	b := 5
	t2 := reflect.TypeOf(b)

	fmt.Println(t2.AssignableTo(t1))
}

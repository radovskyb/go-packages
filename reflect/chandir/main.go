package main

import (
	"fmt"
	"reflect"
)

func main() {
	var chOne chan<- int
	var typeChOne reflect.Type = reflect.TypeOf(chOne)

	// ChanDir returns a channel type's direction.
	// It panics if the type's Kind is not Chan.
	fmt.Println("Channel chOne's type is:", typeChOne.ChanDir().String())
	fmt.Println("Channel chOne - Send direction only:",
		typeChOne.ChanDir() == reflect.SendDir)
	fmt.Println("Channel chOne - Receive direction only:",
		typeChOne.ChanDir() == reflect.RecvDir)

	fmt.Println()

	var chTwo <-chan int
	var typeChTwo reflect.Type = reflect.TypeOf(chTwo)

	fmt.Println("Channel chTwo's type is:", typeChTwo.ChanDir().String())
	fmt.Println("Channel chTwo - Send direction only:",
		typeChTwo.ChanDir() == reflect.SendDir)
	fmt.Println("Channel chTwo - Receive direction only:",
		typeChTwo.ChanDir() == reflect.RecvDir)

	fmt.Println()

	var chThree chan int
	var typeChThree reflect.Type = reflect.TypeOf(chThree)

	fmt.Println("Channel chThree's type is:",
		typeChThree.ChanDir().String())
	fmt.Println("Channel chThree - Send direction only:",
		typeChThree.ChanDir() == reflect.SendDir)
	fmt.Println("Channel chThree - Receive direction only:",
		typeChThree.ChanDir() == reflect.RecvDir)
	fmt.Println("Channel chThree - Both directions (Send & Receive):",
		typeChThree.ChanDir() == reflect.BothDir)
}

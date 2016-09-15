package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create a new channel of type string to use with reflect.Select
	var chr = make(chan string) // Channel Receive (chr)

	// Create a new reflect select case variable rCase
	var rCase reflect.SelectCase

	// Set the rCase.Dir (direction) to SelectRecv
	//
	// Dir is the setting for the direction of the case
	//
	// SelectDir describes the communication direction of a select case.
	// const (
	// 	_             SelectDir = iota
	// 	SelectSend              // case Chan <- Send
	// 	SelectRecv              // case <-Chan:
	// 	SelectDefault           // default
	// )
	rCase.Dir = reflect.SelectRecv // case <-Chan:

	// Set the rCase.Chan (channel) to a new reflect.Value initialized
	// to the value of the channel chr
	rCase.Chan = reflect.ValueOf(chr)

	fmt.Println("Direction of case rCase:", rCase.Dir) // 2 (SelectRecv)
	fmt.Println("Channel to use:", &rCase.Chan, rCase.Chan)

	// Invalid since the channel direction is set to SelectRecv
	fmt.Println("Value to send:", rCase.Send) // Invalid

	fmt.Println()

	var chs = make(chan string, 2) // Channel Send (chs)
	var sCase reflect.SelectCase

	sCase.Dir = reflect.SelectSend    // case Chan <- Send
	sCase.Send = reflect.ValueOf(chs) // value to send (for send)

	fmt.Println("Direction of case sCase:", sCase.Dir) // 1 (SelectSend)

	// Invalid since sCase.Chan is ONLY send, not send or receive
	fmt.Println("Channel to use:", &sCase.Chan, sCase.Chan)

	// Valid because this is the SelectSend case
	fmt.Println("Value to send:", &sCase.Send, sCase.Send)
}

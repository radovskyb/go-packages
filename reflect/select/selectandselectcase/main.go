package main

import (
	"fmt"
	"reflect"
)

func main() {
	var sendCh = make(chan int)

	var increaseInt = func(c chan int) {
		for i := 0; i < 10; i++ {
			sendCh <- i
		}
		close(c)
	}

	go increaseInt(sendCh)

	// A SelectCase describes a single case in a select operation.
	// The kind of case depends on Dir, the communication direction.
	//
	// If Dir is SelectDefault, the case represents a default case.
	// Chan and Send must be zero Values.
	//
	// If Dir is SelectSend, the case represents a send operation.
	// Normally Chan's underlying value must be a channel, and Send's
	// underlying value must be assignable to the channel's element type.
	// As a special case, if Chan is a zero Value, then the case is
	// ignored, and the field Send will also be ignored and may be
	// either zero or non-zero.
	//
	// If Dir is SelectRecv, the case represents a receive operation.
	// Normally Chan's underlying value must be a channel and Send must
	// be a zero Value. If Chan is a zero Value, then the case is
	// ignored, but Send must still be a zero Value.
	// When a receive operation is selected, the received Value is
	// returned by Select.
	var selectCase = make([]reflect.SelectCase, 1)
	selectCase[0].Dir = reflect.SelectRecv
	selectCase[0].Chan = reflect.ValueOf(sendCh)

	var counter int
	for counter < 1 {
		// Select executes a select operation described by the
		// list of cases. Like the Go select statement, it blocks
		// until at least one of the cases can proceed, makes a
		// uniform pseudo-random choice, and then executes that
		// case. It returns the index of the chosen case and, if that
		// case was a receive operation, the value received and a
		// boolean indicating whether the value corresponds to a
		// send on the channel (as opposed to a zero value received
		// because the channel is closed).
		chosen, recv, recvOK := reflect.Select(selectCase)

		if recvOK {
			// Example: Chose case 0, Int was 1, Received was true
			fmt.Println(chosen, recv.Int(), recvOK)
		} else {
			// Loop until the channel is closed, then the counter
			// will increment, therefore breaking out of the loop
			counter++
		}
	}
}

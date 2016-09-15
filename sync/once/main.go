package main

import (
	"fmt"
	"sync"
)

type Object struct {
	// Embed sync.Once's type into Object
	//
	// Once is an object that will perform exactly one action.
	sync.Once
}

// init initializes Object o and prints out a message that it is initializing
func (o *Object) init() {
	fmt.Println("Initializing...")
}

func main() {
	// Create a new Object o
	obj := new(Object)

	// Call obj.Do() on o's init function
	//
	// Do calls the function f if and only if Do is being called for the
	// first time for this instance of Once. In other words, given
	// 	var once Once
	// if once.Do(f) is called multiple times, only the first call will invoke f,
	// even if f has a different value in each invocation.  A new instance of
	// Once is required for each function to execute.
	//
	// Do is intended for initialization that must be run exactly once.  Since f
	// is niladic, it may be necessary to use a function literal to capture the
	// arguments to a function to be invoked by Do:
	// 	config.once.Do(func() { config.init(filename) })
	//
	// Because no call to Do returns until the one call to f returns, if f causes
	// Do to be called, it will deadlock.
	//
	// If f panics, Do considers it to have returned; future calls of Do return
	// without calling f.
	//
	obj.Do(obj.init)

	// Trying to run obj.Do(obj.init) again will not call obj.init
	// as it has already been run
	obj.Do(obj.init)
}

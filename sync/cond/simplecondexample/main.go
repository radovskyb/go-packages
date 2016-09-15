package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func waiter(cond *sync.Cond, id int) {
	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()
	fmt.Printf("Waiter: Wake up waiter %d!\n", id)
	wg.Done()
}

func main() {
	// Create a new mutex to use as `cond`'s locker
	mu := new(sync.Mutex)

	// Create a new `sync.Cond` object by calling `sync.NewCond`
	// and pass it `mu` as it's locker `L`
	//
	// NewCond returns a new Cond with Locker l.
	//
	// Cond implements a condition variable, a rendezvous point
	// for goroutines waiting for or announcing the occurrence
	// of an event.
	//
	// Each Cond has an associated Locker L (often a *Mutex or *RWMutex),
	// which must be held when changing the condition and
	// when calling the Wait method.
	//
	// A Cond can be created as part of other structures.
	// A Cond must not be copied after first use.
	cond := sync.NewCond(mu)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go waiter(cond, i)
	}
	time.Sleep(time.Second)

	// Wake one goroutine waiting on cond by calling cond.Signal
	//
	// Signal wakes one goroutine waiting on c, if there is any.
	//
	// It is allowed but not required for the caller to hold c.L
	// during the call.
	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()

	for i := 3; i < 5; i++ {
		wg.Add(1)
		go waiter(cond, i)
	}
	time.Sleep(time.Second)

	// Signal one more goroutine waiting on cond
	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()

	// Signal all goroutines waiting on cond by calling cond.Broadcast
	//
	// Broadcast wakes all goroutines waiting on c.
	//
	// It is allowed but not required for the caller to hold c.L
	// during the call.
	cond.L.Lock()
	cond.Broadcast()
	cond.L.Unlock()

	wg.Wait()
}

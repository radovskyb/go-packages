package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

func main() {
	// DialHTTP connects to an HTTP RPC server at the specified
	// network address listening on the default HTTP RPC path.
	client, err := rpc.DialHTTP("tcp", ":9000")
	if err != nil {
		log.Fatalln("Dial error: " + err.Error())
	}

	// Create an integer `reply` which will hold the
	// reply from the rpc call
	var reply int

	// Create a new `Args` type called `args` which will hold
	// the arguments to be passed to the rpc server's method
	args := &Args{A: 5, B: 5}

	// Call the method Arith.Add from the rpc server over the rpc client
	//
	// Call invokes the named function, waits for it to complete,
	// and returns its error status.
	err = client.Call("Arith.Add", args, &reply)
	if err != nil {
		log.Fatalln("Arith.Add error: " + err.Error())
	}

	// Print out the result from adding args.A and args.B
	// over the rpc connection by calling Arith.Add
	fmt.Printf("Arith.Add Sync: %d + %d = %d\n", args.A, args.B, reply)

	// Call the Arith.Add method from the rpc server asynchronously as
	// opposed to synchronously like the above method:
	// ---------------------------------------------------------------

	// First change the args.A and args.B integer values
	args.A = 10
	args.B = 10

	// Call the Arith.Add method from the rpc server asynchronously and
	// return an active RPC (remote procedure call) called `call`
	//
	// Go invokes the function asynchronously. It returns the Call
	// structure representing the invocation. The done channel will
	// signal when the call is complete by returning the same Call object.
	// If done is nil, Go will allocate a new channel.
	// If non-nil, done must be buffered or Go will deliberately crash.
	//
	// Call represents an active RPC.
	call := client.Go("Arith.Add", args, &reply, nil)

	// Create a select statement to either wait for the remote
	// procedure call's done channel to be filled, letting us know
	// that the call has been completed, otherwise if there is no
	// response after 10 seconds, timeout and print a message
	select {
	case <-call.Done: // Receive that the remote procedure call is done
		// Print out the reply
		fmt.Printf("Arith.Add Async: %d + %d = %d\n", args.A, args.B, reply)
	case <-time.After(10 * time.Second): // Timeout after 10 seconds
		// Print out a timed out message
		fmt.Println("Timed out after 10 seconds.")
	}
}

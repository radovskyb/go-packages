package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"time"
)

type Args struct {
	A, B int
}

type Arith int

func main() {
	// Create a new jsonrpc client
	//
	// NewClient returns a new rpc.Client to handle requests to the
	// set of services at the other end of the connection.
	//
	// Client will use json as the server's specified codec
	// to encode requests and decode responses.
	// client := jsonrpc.NewClient(conn)

	// Dial connects to a JSON-RPC server at the specified network address.
	client, err := jsonrpc.Dial("tcp", ":9000")
	if err != nil {
		log.Fatalln("Dial error: " + err.Error())
	}

	// Create a new integer `reply`
	var reply int

	// Create a new `Args` type object `args` and
	// set it's `A` and `B` integer values
	args := &Args{A: 10, B: 10}

	// Call the `Arith.Add` method from the rpc server and return
	// a new remote procedure call object `call`
	//
	// Go invokes the function asynchronously. It returns the Call
	// structure representing the invocation. The done channel will signal
	// when the call is complete by returning the same Call object.
	// If done is nil, Go will allocate a new channel.
	// If non-nil, done must be buffered or Go will deliberately crash.
	call := client.Go("Arith.Add", args, &reply, nil)
	if err != nil {
		log.Fatalln("Call error: " + err.Error())
	}

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

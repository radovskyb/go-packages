package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Arith int

func (a *Arith) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func main() {
	// Create a new `Arith` type object `arith`
	arith := new(Arith)

	// Register publishes the receiver's methods in the DefaultServer.
	if err := rpc.Register(arith); err != nil {
		log.Fatalln(err)
	}

	// HandleHTTP registers an HTTP handler for RPC messages to DefaultServer
	// on DefaultRPCPath and a debugging handler on DefaultDebugPath.
	// It is still necessary to invoke http.Serve(), typically
	// in a go statement.
	rpc.HandleHTTP()

	// Create a new net listener on port 9000
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}

	// http.Serve will close the listener when it's done with it
	// so there's no need to call defer ln.Close() above.
	//
	// At this point, clients can see a service "Arith" with the
	// method "Arith.Add"
	//
	// Serve accepts incoming HTTP connections on the listener l,
	// creating a new service goroutine for each.  The service goroutines
	// read requests and then call handler to reply to them.
	// Handler is typically nil, in which case the DefaultServeMux is used.
	log.Fatal(http.Serve(ln, nil))
}

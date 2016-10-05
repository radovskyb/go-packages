package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
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
	arith := new(Arith)

	// Create a new server rather than just using the
	// DefaultServer from rpc when no server is created
	//
	// NewServer returns a new Server.
	server := rpc.NewServer()

	// Register `arith` to the rpc server `server`
	//
	// Register publishes in the server the set of methods of the
	// receiver value that satisfy the following conditions:
	//	- exported method of exported type
	//	- two arguments, both of exported type
	//	- the second argument is a pointer
	//	- one return value, of type error
	// It returns an error if the receiver is not an exported type or has
	// no suitable methods. It also logs the error using package log.
	// The client accesses each method using a string of the form
	// "Type.Method", where Type is the receiver's concrete type.
	if err := server.Register(arith); err != nil {
		log.Fatalln(err)
	}

	// Create a new tcp listener on port 9000
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln("Dial error: " + err.Error())
	}
	defer ln.Close()

	for {
		// Accept a connection from tcp listener `ln`
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln("Accept error: " + err.Error())
		}

		go func(conn net.Conn) {
			defer conn.Close()

			// Call `server.ServeCodec` and call the function
			// `jsonrpc.NewServerCodec(conn)` as it's parameter
			//
			// ServeCodec is like ServeConn but uses the specified codec to
			// decode requests and encode responses.
			//
			// NewServerCodec returns a new rpc.ServerCodec using JSON-RPC on conn.
			server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}(conn)
	}
}

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

// MAX_CLIENTS is only set to 2 for example purposes
// to allow viewing the cond variable at work taking effect
// when new connections are coming in to be created.
//
// This way, only 2 client connections are needed to see sync.Cond
// wait until a client is disconnected before allowing a new
// connection to be created.
const MAX_CLIENTS int = 2

type Server struct {
	cond    *sync.Cond
	clients int
}

func (s *Server) Listen(address string) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Listen error:", err)
	}
	defer ln.Close()

	for {
		s.cond.L.Lock()
		// Wait until client's is less than MAX_CLIENTS
		// before creating another connection
		for s.clients == MAX_CLIENTS {
			// Wait for a signal from cond before checking if the
			// for loop condition is still true
			s.cond.Wait()
		}
		s.cond.L.Unlock()

		// If s.client's is less than MAX_CLIENTS
		// then create a new client connection
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		// After creating the new client connection,
		// increment s.client's
		s.cond.L.Lock()
		s.clients++
		s.cond.L.Unlock()

		// Now finally pass the connection to the s.handleClient
		// method, but spin it off in it's own goroutine so we
		// don't block and can accept more connections asynchronously
		go s.handleClient(conn)
	}
}

func (s *Server) handleClient(conn net.Conn) {
	defer s.disconnected(conn)

	mw := io.MultiWriter(os.Stdout, conn)

	msg := fmt.Sprintf("Welcome Client %d\n", s.clients)

	_, err := mw.Write([]byte(msg))
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *Server) disconnected(conn net.Conn) {
	if err := conn.Close(); err != nil {
		log.Fatalln(err)
	}
	s.cond.L.Lock()
	s.clients--
	s.cond.L.Unlock()
	s.cond.Signal()
}

func main() {
	s := &Server{
		cond: &sync.Cond{L: new(sync.Mutex)},
	}

	s.Listen("localhost:9000")
}

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		defer conn.Close()

		go func() {
			io.Copy(os.Stdout, conn)
		}()
	}
}

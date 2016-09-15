package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		defer conn.Close()

		go func() {
			mw := io.MultiWriter(os.Stdout, conn)
			io.Copy(mw, conn)
		}()
	}
}

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
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()

			mw := io.MultiWriter(os.Stdout, conn)
			_, err = io.Copy(mw, conn)
			if err != nil {
				panic(err)
			}
		}(conn)
	}
}

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := ln.Accept()
		defer conn.Close()
		fmt.Println("Connection accepted.")
		if err != nil {
			log.Fatalln(err)
		}
		go func() {
			io.Copy(os.Stdout, conn)
		}()
	}
}

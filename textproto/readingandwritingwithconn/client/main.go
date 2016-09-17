package main

import (
	"fmt"
	"io"
	"log"
	"net/textproto"
)

func main() {
	// Dial connects to the given address on the given network
	// using net.Dial and then returns a new Conn for the connection.
	conn, err := textproto.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatalln(err)
	}

	// Close closes the connection.
	defer conn.Close()

	// PrintfLine writes the formatted output followed by \r\n.
	if err := conn.PrintfLine("%s", "100 Code 100, From Client!"); err != nil {
		log.Fatalln(err)
	}
	if err := conn.PrintfLine("%s", "200 Code 200, From Client!"); err != nil {
		log.Fatalln(err)
	}
	if err := conn.PrintfLine("%s", "300 Code 300, From Client!"); err != nil {
		log.Fatalln(err)
	}
	if err := conn.PrintfLine("%s", "400 Code 400, From Client!"); err != nil {
		log.Fatalln(err)
	}
	if err := conn.PrintfLine("%s", "500 Code 500, From Client!"); err != nil {
		log.Fatalln(err)
	}

	for {
		// ReadLine reads a single line from r,
		// eliding the final \n or \r\n from the returned string.
		line, err := conn.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Fatalln(err)
			}
		}
		fmt.Println(line)
	}
}

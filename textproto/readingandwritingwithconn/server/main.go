package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/textproto"
	"os"
	"strings"
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

		go handleConnections(conn)
	}
}

func handleConnections(conn net.Conn) {
	defer conn.Close()

	mw := io.MultiWriter(os.Stdout, conn)

	// NewReader returns a new Reader reading from r.
	//
	// To avoid denial of service attacks, the provided bufio.Reader
	// should be reading from an io.LimitReader or similar Reader to bound
	// the size of responses.
	tp := textproto.NewReader(bufio.NewReader(conn))
	code := 100
	for {
		if found, msg := readCode(code, tp); !found {
			continue
		} else {
			_, err := io.Copy(mw, strings.NewReader(msg))
			if err != nil {
				log.Fatalln(err)
			}
		}
		code += 200
	}
}

func readCode(codeInt int, tp *textproto.Reader) (bool, string) {
	// ReadCodeLine reads a response code line of the form
	//	code message
	// where code is a three-digit status code and the message
	// extends to the rest of the line.  An example of such a line is:
	//	220 plan9.bell-labs.com ESMTP
	//
	// If the prefix of the status does not match the digits in expectCode,
	// ReadCodeLine returns with err set to &Error{code, message}.
	// For example, if expectCode is 31, an error will be returned if
	// the status is not in the range [310,319].
	//
	// If the response is multi-line, ReadCodeLine returns an error.
	//
	// An expectCode <= 0 disables the check of the status code.
	//
	code, msg, err := tp.ReadCodeLine(codeInt)
	if err != nil {
		if err != io.EOF {
			return false, ""
		}
	} else {
		msg = fmt.Sprintf("%d - %s\n", code, msg)
	}
	return true, msg
}

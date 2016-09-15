package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/textproto"
)

func main() {
	readBuffer := bytes.NewBuffer([]byte("100 Command Okay.\n"))

	reader := *bufio.NewReader(readBuffer)

	tpReader := textproto.NewReader(&reader)

	// ReadResponse reads a multi-line response of the form:
	//
	//	code-message line 1
	//	code-message line 2
	//	...
	//	code message line n
	//
	// where code is a three-digit status code. The first line starts with the
	// code and a hyphen. The response is terminated by a line that starts
	// with the same code followed by a space. Each line in message is
	// separated by a newline (\n).
	//
	// See page 36 of RFC 959 (http://www.ietf.org/rfc/rfc959.txt) for
	// details.
	//
	// If the prefix of the status does not match the digits in expectCode,
	// ReadResponse returns with err set to &Error{code, message}.
	// For example, if expectCode is 31, an error will be returned if
	// the status is not in the range [310,319].
	//
	// An expectCode <= 0 disables the check of the status code.
	code, msg, err := tpReader.ReadResponse(100)
	if err != nil {
		fmt.Println("Read error:", err)
	}

	fmt.Println("Code:", code)
	fmt.Println("Message:", msg)
}

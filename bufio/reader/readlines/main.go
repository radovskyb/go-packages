package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World!\nWat's up doc?\nLadee Da Dee Dah!"

	// Create a new buffered reader br which reads from
	// a new string reader that reads from the string s
	br := bufio.NewReader(strings.NewReader(s))

	// Read one line froma br
	//
	// ReadLine tries to return a single line, not including the end-of-line bytes.
	// If the line was too long for the buffer then isPrefix is set and the beginning
	// of the line is returned. The rest of the line will be returned from future
	// calls. isPrefix will be false when returning the last fragment of the line.
	// The returned buffer is only valid until the next call to ReadLine. ReadLine
	// either returns a non-nil line or it returns an error, never both.
	//
	// The text returned from ReadLine does not include the line end ("\r\n" or "\n").
	// No indication or error is given if the input ends without a final line end.
	// Calling UnreadByte after ReadLine will always unread the last byte read (possibly
	// a character belonging to the line end) even if that byte is not part of the
	// line returned by ReadLine.
	//
	// ReadLine is a low-level line-reading primitive. Most callers should use
	// ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
	line, _, err := br.ReadLine()
	if err != nil {
		log.Fatalln(err)
	}

	// Print the line read from br in string form
	fmt.Println(string(line)) // Hello, World
}

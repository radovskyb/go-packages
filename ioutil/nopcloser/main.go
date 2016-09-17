package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	str := bytes.NewReader([]byte("abcdefg"))

	// Create a new read closer with a no-op close method
	// by using ioutil.NopCloser
	//
	// NopCloser returns a ReadCloser with a no-op Close method
	// wrapping the provided Reader r.
	//
	// One reason to use NopCloser is as follows:
	// Whenever you need to return an io.ReadCloser, while making
	// sure a Close() function is available, you can use a NopCloser
	// to build such a ReaderCloser.
	reader := ioutil.NopCloser(str)

	// We can now close the Read Closer returned from NopCloser
	defer reader.Close()

	// Create a new byte slice buf of size 7 bytes
	buf := make([]byte, 7)

	// Use reader to read it's contents into buf
	_, err := reader.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}

	// Print out buf in string form
	fmt.Println(string(buf))
}

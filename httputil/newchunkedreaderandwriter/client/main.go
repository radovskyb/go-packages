package main

import (
	"io"
	"log"
	"net"
	"net/http/httputil"
	"os"
)

func main() {
	var chunks = "4\r\nWiki\r\n5\r\npedia\r\ne\r\n in\r\n\r\nchunks.\r\n0\r\n\r\n"

	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// NewChunkedWriter returns a new chunkedWriter that translates writes into HTTP
	// "chunked" format before writing them to w. Closing the returned chunkedWriter
	// sends the final 0-length chunk that marks the end of the stream.
	//
	// NewChunkedWriter is not needed by normal applications. The http
	// package adds chunking automatically if handlers don't set a
	// Content-Length header. Using NewChunkedWriter inside a handler
	// would result in double chunking or chunking with a Content-Length
	// length, both of which are wrong.
	cw := httputil.NewChunkedWriter(conn)
	_, err = cw.Write([]byte(chunks))
	if err != nil {
		log.Fatalln(err)
	}

	// NewChunkedReader returns a new chunkedReader that translates the data read from r
	// out of HTTP "chunked" format before returning it.
	// The chunkedReader returns io.EOF when the final 0-length chunk is read.
	//
	// NewChunkedReader is not needed by normal applications. The http package
	// automatically decodes chunking when reading response bodies.
	cr := httputil.NewChunkedReader(conn)
	_, err = io.Copy(os.Stdout, cr)
	if err != nil {
		log.Fatalln(err)
	}
}

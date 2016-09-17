package main

import (
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	// A PipeReader is the read half of a pipe.
	// A PipeWriter is the write half of a pipe.
	//
	// Pipe creates a synchronous in-memory pipe. It can be used to
	// connect code expecting an io.Reader with code expecting an
	// io.Writer. Reads on one end are matched with writes on the other,
	// copying data directly between the two; there is no internal buffering.
	// It is safe to call Read and Write in parallel with each other or
	// with Close. Close will complete once pending I/O is done. Parallel
	// calls to Read, and parallel calls to Write, are also safe:
	// the individual calls will be gated sequentially.
	pReader, pWriter := io.Pipe()

	// use base64 encoder
	b64Writer := base64.NewEncoder(base64.StdEncoding, pWriter)

	gWriter := gzip.NewWriter(b64Writer)

	// write text to be gzipped and push to pipe
	go func() {
		fmt.Println("Start writing")

		n, err := gWriter.Write([]byte("These words will be compressed and push into pipe"))
		if err != nil {
			log.Fatalln(err)
		}

		if err := gWriter.Close(); err != nil {
			log.Fatalln(err)
		}
		if err := b64Writer.Close(); err != nil {
			log.Fatalln(err)
		}
		if err := pWriter.Close(); err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Written %d bytes \n", n)
	}()

	// start reading from the pipe(reverse of the above process)

	// use base64 decoder to wrap pipe Reader
	b64Reader := base64.NewDecoder(base64.StdEncoding, pReader)

	// read gzipped text and decompressed the text
	gReader, err := gzip.NewReader(b64Reader)

	if err != nil {
		log.Fatalln(err)
	}

	// print out the text
	text, err := ioutil.ReadAll(gReader)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", text)
}

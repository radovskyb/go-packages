package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var b bytes.Buffer

	// NewWriter creates a new Writer.
	// Writes to the returned Writer are compressed and written to w.
	//
	// It is the caller's responsibility to call Close on the
	// WriteCloser when done.
	// Writes may be buffered and not flushed until Close.
	w := zlib.NewWriter(&b)

	// Write writes a compressed form of p to the underlying io.Writer. The
	// compressed bytes are not necessarily flushed until the Writer is closed or
	// explicitly flushed.
	_, err := w.Write([]byte("Hello, World!\n"))
	if err != nil {
		log.Fatalln(err)
	}

	// Close closes the Writer, flushing any unwritten data to the underlying
	// io.Writer, but does not close the underlying io.Writer.
	if err := w.Close(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(b.String())

	// NewReader creates a new ReadCloser.
	// Reads from the returned ReadCloser read and decompress data from r.
	// The implementation buffers input and may read more data than
	// necessary from r.
	// It is the caller's responsibility to call Close on the ReadCloser
	// when done.
	//
	// The ReadCloser returned by NewReader also implements Resetter.
	r, err := zlib.NewReader(&b)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()

	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}

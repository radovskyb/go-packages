package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	// Import quotedprintable as qp since it's a long name to type.
	qp "mime/quotedprintable"
	"os"
)

func main() {
	var buf bytes.Buffer

	// NewWriter returns a new Writer that writes to w.
	w := qp.NewWriter(&buf)

	// Write encodes p using quoted-printable encoding and writes it to the
	// underlying io.Writer. It limits line length to 76 characters. The encoded
	// bytes are not necessarily flushed until the Writer is closed.
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		log.Fatalln(err)
	}

	// Binary mode treats the writer's input as pure binary and processes end of
	// line bytes as binary data.
	fmt.Println(w.Binary)

	// Close closes the Writer, flushing any unwritten data to the underlying
	// io.Writer, but does not close the underlying io.Writer.
	if err := w.Close(); err != nil {
		log.Fatalln(err)
	}

	// NewReader returns a quoted-printable reader, decoding from r.
	r := qp.NewReader(&buf)

	// Read reads and decodes quoted-printable data from the underlying reader.
	_, err = io.Copy(os.Stdout, r) // Calls the r.Read method underneath
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"bytes"
	"fmt"
	"io"
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
	w.Write([]byte("Hello, World!"))

	// Binary mode treats the writer's input as pure binary and processes end of
	// line bytes as binary data.
	fmt.Println(w.Binary)

	// Close closes the Writer, flushing any unwritten data to the underlying
	// io.Writer, but does not close the underlying io.Writer.
	w.Close()

	// NewReader returns a quoted-printable reader, decoding from r.
	r := qp.NewReader(&buf)

	// Read reads and decodes quoted-printable data from the underlying reader.
	io.Copy(os.Stdout, r) // Calls the r.Read method underneath
}

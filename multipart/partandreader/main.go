package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/mail"
	"strings"
)

func main() {
	msg := &mail.Message{
		Header: map[string][]string{
			"Content-Type": {"multipart/mixed; boundary=foo"},
		},
		Body: strings.NewReader(
			"--foo\r\nFoo: one\r\n\r\nA section\r\n" +
				"--foo\r\nFoo: two\r\n\r\nAnd another\r\n" +
				"--foo--\r\n"),
	}

	mediaType, params, err := mime.ParseMediaType(msg.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		// NewReader creates a new multipart Reader reading from r using the
		// given MIME boundary.
		//
		// The boundary is usually obtained from the "boundary" parameter of
		// the message's "Content-Type" header. Use mime.ParseMediaType to
		// parse such headers.
		mr := multipart.NewReader(msg.Body, params["boundary"])

		for {
			// NextPart returns the next part in the multipart or an error.
			// When there are no more parts, the error io.EOF is returned.
			p, err := mr.NextPart()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatal(err)
			}

			slurp, err := ioutil.ReadAll(p)
			if err != nil {
				log.Fatal(err)
			}

			// FormName returns the name parameter if p has a Content-Disposition
			// of type "form-data". Otherwise it returns the empty string.
			if formname := p.FormName(); formname != "" {
				fmt.Println(formname + ":")
			}

			// FileName returns the filename parameter of the Part's
			// Content-Disposition header.
			if filename := p.FileName(); filename != "" {
				fmt.Println(filename + ":")
			}

			fmt.Printf("Part %q: %q\n", p.Header.Get("Foo"), slurp)

			// Close the part
			if err := p.Close(); err != nil {
				log.Fatalln(err)
			}
		}
	}
}

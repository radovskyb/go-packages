package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// A reader that is reading from a file or any thing else that can
	// be read from could be used instead of a simple strings reader
	rdr := strings.NewReader("Hello, World from http.ServeContent()")

	// ServeContent replies to the request using the content in the
	// provided ReadSeeker.  The main benefit of ServeContent over io.Copy
	// is that it handles Range requests properly, sets the MIME type, and
	// handles If-Modified-Since requests.
	//
	// If the response's Content-Type header is not set, ServeContent
	// first tries to deduce the type from name's file extension and,
	// if that fails, falls back to reading the first block of the content
	// and passing it to DetectContentType.
	// The name is otherwise unused; in particular it can be empty and is
	// never sent in the response.
	//
	// If modtime is not the zero time or Unix epoch, ServeContent
	// includes it in a Last-Modified header in the response.  If the
	// request includes an If-Modified-Since header, ServeContent uses
	// modtime to decide whether the content needs to be sent at all.
	//
	// The content's Seek method must work: ServeContent uses
	// a seek to the end of the content to determine its size.
	//
	// If the caller has set w's ETag header, ServeContent uses it to
	// handle requests using If-Range and If-None-Match.
	//
	// Note that *os.File implements the io.ReadSeeker interface.
	http.ServeContent(w, r, "handler", time.Now(), rdr)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

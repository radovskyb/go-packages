package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	handler := &homeHandler{}

	// NewServer starts and returns a new Server.
	// The caller should call Close when finished, to shut it down.
	server := httptest.NewServer(handler)

	// Close shuts down the server and blocks until all outstanding
	// requests on this server have completed.
	defer server.Close()

	for _, i := range []int{1, 2, 3, 4, 5} {
		resp, err := http.Get(server.URL)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != 200 {
			t.Fatalf("Received non-200 response: %d", resp.StatusCode)
		}

		expected := fmt.Sprintf("Visitor count: %d.", i)
		actual, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if err := resp.Body.Close(); err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(actual), expected) {
			t.Errorf("expected body to contain %s\nbody: %s", expected, actual)
		}
	}
}

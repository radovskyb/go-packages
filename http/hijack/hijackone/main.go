package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	conn, _, _ := w.(http.Hijacker).Hijack()
	conn.Write([]byte("Server has hijacked this connection."))
	conn.Close()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}

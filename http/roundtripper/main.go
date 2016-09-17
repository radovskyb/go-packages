package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	// DefaultTransport is the default implementation of Transport and is
	// used by DefaultClient. It establishes network connections as needed
	// and caches them for reuse by subsequent calls. It uses HTTP proxies
	// as directed by the $HTTP_PROXY and $NO_PROXY (or $http_proxy and
	// $no_proxy) environment variables.
	// var DefaultTransport RoundTripper = &Transport{
	// 	Proxy: ProxyFromEnvironment,
	// 	Dial: (&net.Dialer{
	// 		Timeout:   30 * time.Second,
	// 		KeepAlive: 30 * time.Second,
	// 	}).Dial,
	// 	TLSHandshakeTimeout: 10 * time.Second,
	// }
	req, err := http.NewRequest("HEAD", "http://betsee.com.au", nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	if err := resp.Body.Close(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp)
}

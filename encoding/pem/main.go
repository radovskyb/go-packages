package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	pemfile, err := os.Create("cert.pem")
	if err != nil {
		log.Fatalln("Could not create file cert.pem:", err)
	}

	// Create a new rsa.PrivateKey object to use with pem.Block
	// as it's `Bytes:` field
	var privatekey *rsa.PrivateKey

	// Generate an rsa.PrivateKey use rsa.Generate of size 1024 bytes
	privatekey, err = rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Fatalln(err)
	}

	// Create a new pem Block object to use with pem.Encode
	//
	// A Block represents a PEM encoded structure.
	//
	// The encoded form is:
	//    -----BEGIN Type-----
	//    Headers
	//    base64-encoded Bytes
	//    -----END Type-----
	// where Headers is a possibly empty sequence of Key: Value lines.
	var pemkey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privatekey),
	}

	// Encode pemkey to in to memory and store it in
	// variable inMemoryPrivateKey using pem.EncodeToMemory
	inMemoryPrivateKey := pem.EncodeToMemory(pemkey)

	// Print out inMemoryPrivateKey as a string
	fmt.Println(string(inMemoryPrivateKey))

	fmt.Println()

	// Encode pemkey as a pem dataencoding into pemfile file
	if err := pem.Encode(pemfile, pemkey); err != nil {
		log.Fatalln(err)
	}

	// Close the file handle to pemfile
	if err := pemfile.Close(); err != nil {
		log.Fatalln(err)
	}

	// Now that the file has been written and the file
	// handle is closed, reopen the file and read all of
	// it's contents to print to the screen
	pemfile, err = os.Open("cert.pem")
	if err != nil {
		log.Fatalln("Could not open cert.pem:", err)
	}
	defer pemfile.Close()

	// Read the contents into contents variable
	contents, err := ioutil.ReadAll(pemfile)
	if err != nil {
		log.Fatalln("Could not read cert.pem's contents:", err)
	}

	// Print the contents as a string
	fmt.Println(string(contents))

	// Decode contents into a pem.Block object
	//
	// Decode will find the next PEM formatted block (certificate,
	// private key etc) in the input. It returns that block and the
	// remainder of the input. If no PEM data is found, p is nil and
	// the whole of the input is returned in rest.
	pemBlock, rest := pem.Decode([]byte(contents))

	// In this case rest will be nil, since the file only contains
	// a pem block and no other data before or after it
	fmt.Println(rest)

	fmt.Println("Headers:", pemBlock.Headers)
	fmt.Println("Type:", pemBlock.Type)
	fmt.Println("Bytes:", pemBlock.Bytes)
}

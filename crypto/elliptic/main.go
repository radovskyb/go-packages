package main

import (
	"crypto"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
)

type ellipticPublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}

type ellipticPrivateKey struct {
	privk []byte
}

func main() {
	var privatekey crypto.PrivateKey
	var publickey crypto.PublicKey
	var err error
	var priv []byte
	var x, y *big.Int

	// Using elliptic.P224 for this example
	//
	// P224 returns a Curve which implements P-224
	// (see FIPS 186-3, section D.2.2)
	//
	// GenerateKey returns a public/private key pair. The private key is
	// generated using the given reader, which must return random data.
	priv, x, y, err = elliptic.GenerateKey(elliptic.P224(), rand.Reader)
	if err != nil {
		fmt.Println(err)
	}

	privatekey = &ellipticPrivateKey{
		privk: priv,
	}

	publickey = &ellipticPublicKey{
		Curve: elliptic.P224(),
		X:     x,
		Y:     y,
	}

	fmt.Printf("Private key: %x\n", privatekey)

	fmt.Printf("\nPublic key: %x\n", publickey)
}

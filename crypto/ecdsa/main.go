package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"io"
	"log"
	"math/big"
)

func main() {
	//see http://golang.org/pkg/crypto/elliptic/#P256
	pubkeyCurve := elliptic.P256()

	// PrivateKey represents a ECDSA private key.
	privatekey := new(ecdsa.PrivateKey)

	// this generates a public & private key pair
	//
	// GenerateKey generates a public and private key pair.
	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
	if err != nil {
		log.Fatalln(err)
	}

	// PublicKey represents an ECDSA public key.
	var pubkey ecdsa.PublicKey
	pubkey = privatekey.PublicKey

	fmt.Println("Private Key:")
	fmt.Printf("%x \n", privatekey)

	fmt.Println("\nPublic Key:")
	fmt.Printf("%x \n", pubkey)

	// Sign ecdsa style

	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	_, err = io.WriteString(h, "This is a message to be signed and verified by ECDSA!")
	if err != nil {
		log.Fatalln(err)
	}
	signhash := h.Sum(nil)

	// Sign signs an arbitrary length hash (which should be the result
	// of hashing a larger message) using the private key, priv.
	// It returns the signature as a pair of integers. The security
	// of the private key depends on the entropy of rand.
	r, s, serr := ecdsa.Sign(rand.Reader, privatekey, signhash)
	if serr != nil {
		log.Fatalln(serr)
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	fmt.Printf("\nSignature: %x\n", signature)

	// Verify

	// Verify verifies the signature in r, s of hash using the public
	// key, pub. Its return value records whether the signature is valid.
	verifystatus := ecdsa.Verify(&pubkey, signhash, r, s)
	fmt.Printf("\n`verifystatus`: %t", verifystatus) // true
}

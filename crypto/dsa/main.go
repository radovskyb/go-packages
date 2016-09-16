package main

import (
	"crypto/dsa"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"io"
	"log"
	"math/big"
)

func main() {
	// Parameters represents the domain parameters for a key.
	// These parameters can be shared across many keys. The bit
	// length of Q must be a multiple of 8.
	params := new(dsa.Parameters)

	// Generate and put a random, valid set of DSA parameters into params.
	//
	// see http://golang.org/pkg/crypto/dsa/#ParameterSizes
	//
	// GenerateParameters puts a random, valid set of DSA parameters into params.
	// This function takes many seconds, even on fast machines.
	if err := dsa.GenerateParameters(params, rand.Reader, dsa.L1024N160); err != nil {
		log.Fatalln("GenerateParameters error:", err)
	}

	// Create a new DSA private key
	//
	// PrivateKey represents a DSA private key.
	privatekey := new(dsa.PrivateKey)

	// Set privatekey's publickey's parameters to *params
	privatekey.PublicKey.Parameters = *params

	// Generate a public and private key pair
	//
	// GenerateKey generates a public&private key pair. The Parameters of the
	// PrivateKey must already be valid (see GenerateParameters).
	if err := dsa.GenerateKey(privatekey, rand.Reader); err != nil {
		log.Fatalln(err)
	}

	// Create a new DSA public key variable `publickey`
	//
	// PublicKey represents a DSA public key.
	var publickey dsa.PublicKey

	// set `publickey` to privatekey's public key
	publickey = privatekey.PublicKey

	// Print out the private and public key's
	fmt.Printf("Private Key: %x\n", privatekey)
	fmt.Printf("\nPublic Key: %x\n", publickey)

	// Sign:

	// Create a new hash.Hash variable `h`
	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	_, err := io.WriteString(h, "This is the message to be signed and verified!")
	if err != nil {
		log.Fatalln(err)
	}
	signhash := h.Sum(nil)

	// Sign signs an arbitrary length hash (which should be the result
	// of hashing a larger message) using the private key, priv.
	// It returns the signature as a pair of integers. The security
	// of the private key depends on the entropy of rand.
	//
	// Note that FIPS 186-3 section 4.6 specifies that the hash should
	// be truncated to the byte-length of the subgroup. This function
	// does not perform that truncation itself.
	r, s, err = dsa.Sign(rand.Reader, privatekey, signhash)
	if err != nil {
		log.Fatalln("Sign error:", err)
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	fmt.Printf("\nSignature: %x\n", signature)

	// Verify:

	// Verify verifies the signature in r, s of hash using the
	// public key, pub. It reports whether the signature is valid.
	//
	// Note that FIPS 186-3 section 4.6 specifies that the hash should
	// be truncated to the byte-length of the subgroup. This function
	// does not perform that truncation itself.
	verifystatus := dsa.Verify(&publickey, signhash, r, s)
	fmt.Printf("\n`verifystatus`: %t\n", verifystatus) // true

	// we add additional data to change the signhash
	_, err = io.WriteString(h, "This message is NOT to be signed and verified!")
	if err != nil {
		log.Fatalln(err)
	}
	signhash = h.Sum(nil)

	verifystatus = dsa.Verify(&publickey, signhash, r, s)
	fmt.Printf("\n`verifystatus`: %t\n", verifystatus) // false
}

// An HMAC is a cryptographic hash that uses a key to sign a message.
// The receiver verifies the hash by recomputing it using the same key.
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"log"
)

func ComputeHMAC256(message string, secret string) ([]byte, error) {
	key := []byte(secret)

	// New returns a new HMAC hash using the given hash.Hash type and key.
	h := hmac.New(sha256.New, key)

	// Call `h`'s Write method with `message`
	_, err := h.Write([]byte(message))
	if err != nil {
		return nil, err
	}

	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	return h.Sum(nil), nil
}

func CheckMAC(message string, messageMAC []byte, secret string) (bool, error) {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	_, err := h.Write([]byte(message))
	if err != nil {
		return false, err
	}
	expectedMAC := h.Sum(nil)

	// Equal compares two MACs for equality without leaking timing information.
	return hmac.Equal([]byte(messageMAC), []byte(expectedMAC)), nil
}

func main() {
	// Generate HMAC
	mac, err := ComputeHMAC256("Message", "secret")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%x\n", mac)

	// Check if hmac's are equal
	equal, err := CheckMAC("Message123", mac, "secret")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(equal)

	equal, err = CheckMAC("Message", mac, "secret")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(equal)
}

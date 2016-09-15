// An HMAC is a cryptographic hash that uses a key to sign a message.
// The receiver verifies the hash by recomputing it using the same key.
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func ComputeHMAC256(message string, secret string) []byte {
	key := []byte(secret)

	// New returns a new HMAC hash using the given hash.Hash type and key.
	h := hmac.New(sha256.New, key)

	// Call `h`'s Write method with `message`
	h.Write([]byte(message))

	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	return h.Sum(nil)
}

func CheckMAC(message string, messageMAC []byte, secret string) bool {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	expectedMAC := h.Sum(nil)

	// Equal compares two MACs for equality without leaking timing information.
	return hmac.Equal([]byte(messageMAC), []byte(expectedMAC))
}

func main() {
	// Generate HMAC
	mac := ComputeHMAC256("Message", "secret")
	fmt.Printf("%x\n", mac)

	// Check if hmac's are equal
	fmt.Println(CheckMAC("Message123", mac, "secret"))
	fmt.Println(CheckMAC("Message", mac, "secret"))
}

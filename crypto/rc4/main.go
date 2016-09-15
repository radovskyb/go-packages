package main

import (
	"crypto/rc4"
	"fmt"
)

func main() {
	// Create a new key that can be between 1 to 256 bytes
	key := []byte("123") // min 1 byte, max is 256 bytes

	// Create a new Cipher
	//
	// NewCipher creates and returns a new Cipher. The key argument
	// should be the RC4 key, at least 1 byte and at most 256 bytes.
	c, err := rc4.NewCipher(key)
	if err != nil {
		// this is the KeySizeError output if the key is less than 1 byte
		// or more than 256 bytes
		fmt.Println(err.Error) // KeySizeError
	}

	// Create a `plaintext` []byte that contains the message to encrypt
	plaintext := []byte("Secret message")

	// Encrypt:

	// Create a []byte `encrypted` to store the encrypted `plaintext` into
	encrypted := make([]byte, len(plaintext))

	// Encrypt `plaintext` into `encrypted`
	//
	// XORKeyStream sets dst to the result of XORing src with the key stream.
	// Dst and src may be the same slice but otherwise should not overlap.
	c.XORKeyStream(encrypted, plaintext)

	// Print out the encrypted rc4 `plaintext`
	fmt.Printf("[%s] encrypted to [%x] by rc4 crypto\n", plaintext, encrypted)

	// Reset zeros the key data so that it will no
	// longer appear in the process's memory.
	c.Reset() // Reset the key data just for fun

	// Decrypt:

	// Create a []byte `decrypted` to store the decrypted `encrypted` into
	decrypted := make([]byte, len(encrypted))

	// Regenerate the key after resetting it above just for the sake
	// of calling reset on the cipher
	c, err = rc4.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error)
	}

	// Decrypt `encrypted` into `decrypted`
	c.XORKeyStream(decrypted, encrypted)

	// Print out the decrypted rc4 encrypted data `encrypted`
	fmt.Printf("[%x] decrypted to [%s] \n", encrypted, decrypted)
}

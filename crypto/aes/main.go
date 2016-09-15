package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
)

func main() {
	// The key argument should be the AES key, either 16, 24,
	// or 32 bytes to select AES-128, AES-192, or AES-256.
	key := "opensesame123456"

	// Create a new cipher block from `key`
	//
	// NewCipher creates and returns a new cipher.Block.
	// The key argument should be the AES key,
	// either 16, 24, or 32 bytes to select
	// AES-128, AES-192, or AES-256.
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalln("NewCipher error:", err)
	}

	fmt.Printf("%d bytes NewCipher key with block size of %d bytes\n",
		len(key), block.BlockSize)

	// Create a string to be encrypted
	str := []byte("Hello World!")

	// 16 bytes for AES-128, 24 bytes for AES-192, 32 bytes for AES-256
	ciphertext := []byte("abcdef1234567890")

	iv := ciphertext[:aes.BlockSize] // const BlockSize = 16

	// Encrypt:

	// NewCFBEncrypter returns a Stream which encrypts with cipher feedback mode,
	// using the given Block. The iv must be the same length as the Block's block
	// size.
	encrypter := cipher.NewCFBEncrypter(block, iv)

	encrypted := make([]byte, len(str))

	// XORKeyStream XORs each byte in the given slice with a byte from the
	// cipher's key stream. Dst and src may point to the same memory.
	// If len(dst) < len(src), XORKeyStream should panic. It is acceptable
	// to pass a dst bigger than src, and in that case, XORKeyStream will
	// only update dst[:len(src)] and will not touch the rest of dst.
	encrypter.XORKeyStream(encrypted, str)

	fmt.Printf("%s encrypted to %v\n", str, encrypted)

	// Decrypt:

	// NewCFBDecrypter returns a Stream which decrypts with cipher feedback mode,
	// using the given Block. The iv must be the same length as the Block's block
	// size.
	decrypter := cipher.NewCFBDecrypter(block, iv)

	decrypted := make([]byte, len(str))

	decrypter.XORKeyStream(decrypted, encrypted)

	fmt.Printf("%v decrypts to %s\n", encrypted, decrypted)
}

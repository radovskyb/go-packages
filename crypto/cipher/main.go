package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

func main() {
	// Create a key that's 16 bytes long
	key := []byte("example key 1234")
	plaintext := []byte("exampleplaintext")

	// CBC mode works on blocks so plaintexts may need to be padded to
	// the next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	//
	// BlockSize: The AES block size in bytes. (16 bytes)
	if len(plaintext)%aes.BlockSize != 0 {
		log.Fatalln("`plaintext` is not a multiple of the block size")
	}

	// Create a new cipher block based on the aes key `key`
	//
	// NewCipher creates and returns a new cipher.Block.
	// The key argument should be the AES key,
	// either 16, 24, or 32 bytes to select
	// AES-128, AES-192, or AES-256.
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln("NewCipher error:", err)
	}

	// Encrypt:

	// Create a byte slice to store the encrypted data into
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// The IV needs to be unique, but not secure. Therefore it's
	// common to include it at the beginning of the ciphertext.
	//
	// An initialization vector (IV) is an arbitrary number that can
	// be used along with a secret key for data encryption. This number,
	// also called a nonce, is employed only one time in any session.
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatalln("ReadFull error:", err)
	}

	// Set the block mode
	//
	// NewCBCEncrypter returns a BlockMode which encrypts in cipher
	// block chaining mode, using the given Block. The length of iv
	// must be the same as the Block's block size.
	mode := cipher.NewCBCEncrypter(block, iv)

	// CryptBlocks encrypts or decrypts a number of blocks. The length of
	// src must be a multiple of the block size. Dst and src may point to
	// the same memory.
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	// Encryped text
	fmt.Printf("%s encrypts to:\n%x\n", plaintext, ciphertext)

	// Store a temporary string of what the encrypted string was
	// to print out after CryptoBlocks decrypts ciphertext in it's
	// same memory location when we are doing decryption since
	// you can decrypt an encrypted string into it's own memory location
	encrypted := fmt.Sprintf("%x", ciphertext)

	fmt.Println()

	// Decrypt:

	if len(ciphertext) < aes.BlockSize {
		log.Fatalln("`ciphertext` is too short")
	}

	iv = ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		log.Fatalln("`ciphertext` is not a multiple of the block size")
	}

	// Set the block mode
	//
	// NewCBCDecrypter returns a BlockMode which decrypts in cipher block chaining
	// mode, using the given Block. The length of iv must be the same as the
	// Block's block size and must match the iv used to encrypt the data.
	mode = cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	// Print out the decrypted ciphertext as a string
	fmt.Printf("%s decrypts to:\n%s", encrypted, string(ciphertext))
}

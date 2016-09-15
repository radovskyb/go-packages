package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

func main() {
	// Create a new key
	key := "hello123" // 8 bytes for des.BlockSize

	// Create some plain text to get encryped
	plaintext := "exampleplaintext"

	// BlockSize: The DES block size in bytes.
	if len(plaintext)%des.BlockSize != 0 {
		log.Fatalln("`plaintext` is not a multiple of the block size")
	}

	// Create a new cipher block using `key`
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		log.Fatalln("NewCipher error:", err)
	}

	// Encrypt

	// Create a byte slice to store the encrypted data into
	ciphertext := make([]byte, des.BlockSize+len(plaintext))

	// Create a new initialization vector (IV)
	iv := ciphertext[:des.BlockSize]

	// Read len(iv) random bytes into iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatalln("ReadFull error:", err)
	}

	// Encrypt:

	// Set the block mode
	mode := cipher.NewCBCEncrypter(block, iv)

	// Encrypt `plaintext` and store the encrypted data in `ciphertext`
	mode.CryptBlocks(ciphertext[des.BlockSize:], []byte(plaintext))

	// Print out the encrypted data stored in `ciphertext`
	fmt.Printf("%s encrypts to:\n%x\n", plaintext, ciphertext)

	fmt.Println()

	// Store the encryped data from ciphertext in a variable to
	// use for printing after decrypting since we will decrypting
	// ciphertext into it's own memory location
	encrypted := fmt.Sprintf("%x", ciphertext)

	// Decrypt:

	if len(ciphertext) < des.BlockSize {
		log.Fatalln("`ciphertext` is too short")
	}

	// Set the iv and ciphertext
	iv = ciphertext[:des.BlockSize]
	ciphertext = ciphertext[des.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%des.BlockSize != 0 {
		log.Fatalln("`ciphertext` is not a multiple of the block size")
	}

	// Set the block mode
	mode = cipher.NewCBCDecrypter(block, iv)

	// Decrypt ciphertext and store the decrypted data in it's own
	// memory location
	//
	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// Print out the decrypted results
	fmt.Printf("%s decrypts to:\n%s\n", encrypted, ciphertext)
}

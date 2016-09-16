package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// generate private key
	privatekey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Fatalln(err)
	}

	D := privatekey.D // private exponent
	Primes := privatekey.Primes
	PCValues := privatekey.Precomputed

	// Note : Only used for 3rd and subsequent primes
	// CRTVal := privatekey.Precomputed.CRTValues

	fmt.Println("Private Key : ", privatekey)
	fmt.Println()
	fmt.Println("Private Exponent : ", D.String())
	fmt.Println()
	fmt.Printf("Primes : %s %s \n", Primes[0].String(), Primes[1].String())
	fmt.Println()
	fmt.Printf("Precomputed Values : Dp[%s] Dq[%s]\n", PCValues.Dp.String(), PCValues.Dq.String())
	fmt.Println()
	fmt.Printf("Precomputed Values : Qinv[%s]", PCValues.Qinv.String())
	fmt.Println()

	// Note : Only used for 3rd and subsequent primes
	// fmt.Printf("CRTValues : Exp[%s]\n Coeff[%s]\n R[%s]\n", CRTVal[2].Exp.String(), CRTVal[2].Coeff.String(), CRTVal[2].R.String())

	// Note : if you want to have multi primes,
	//        use rsa.GenerateMultiPrimeKey() function instead of
	//        rsa.GenerateKey() function
	//        see http://golang.org/pkg/crypto/rsa/#GenerateMultiPrimeKey

	// http://golang.org/pkg/crypto/rsa/#PrivateKey.Precompute
	privatekey.Precompute()

	// http://golang.org/pkg/crypto/rsa/#PrivateKey.Validate
	err = privatekey.Validate()
	if err != nil {
		log.Fatalln(err)
	}

	var publickey *rsa.PublicKey
	publickey = &privatekey.PublicKey
	N := publickey.N // modulus
	E := publickey.E // public exponent

	fmt.Println("Public key ", publickey)
	fmt.Println()
	fmt.Println("Public Exponent : ", E)
	fmt.Println()
	fmt.Println("Modulus : ", N.String())
	fmt.Println()

	// EncryptOAEP
	msg := []byte("The secret message!")
	label := []byte("")
	md5hash := md5.New()

	encryptedmsg, err := rsa.EncryptOAEP(md5hash, rand.Reader, publickey, msg, label)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("OAEP encrypted [%s] to \n[%x]\n", string(msg), encryptedmsg)
	fmt.Println()

	// DecryptOAEP
	decryptedmsg, err := rsa.DecryptOAEP(md5hash, rand.Reader, privatekey, encryptedmsg, label)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("OAEP decrypted [%x] to \n[%s]\n", encryptedmsg, decryptedmsg)
	fmt.Println()

	// EncryptPKCS1v15
	encryptedPKCS1v15, errPKCS1v15 := rsa.EncryptPKCS1v15(rand.Reader, publickey, msg)
	if errPKCS1v15 != nil {
		log.Fatalln(err)
	}

	fmt.Printf("PKCS1v15 encrypted [%s] to \n[%x]\n", string(msg), encryptedPKCS1v15)
	fmt.Println()

	// DecryptPKCS1v15
	decryptedPKCS1v15, err := rsa.DecryptPKCS1v15(rand.Reader, privatekey, encryptedPKCS1v15)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("PKCS1v15 decrypted [%x] to \n[%s]\n", encryptedPKCS1v15, decryptedPKCS1v15)
	fmt.Println()

	// SignPKCS1v15
	var h crypto.Hash
	message := []byte("This is the message to be signed!")
	hash := md5.New()
	_, err = io.WriteString(hash, string(message))
	if err != nil {
		log.Fatalln(err)
	}
	hashed := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privatekey, h, hashed)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("PKCS1v15 Signature : %x\n", signature)

	// VerifyPKCS1v15
	err = rsa.VerifyPKCS1v15(publickey, h, hashed, signature)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("VerifyPKCS1v15 successful")
	}
	fmt.Println()

	// SignPSS
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example

	PSSmessage := []byte("Message to be PSSed!")
	newhash := crypto.MD5
	pssh := newhash.New()
	_, err = pssh.Write(PSSmessage)
	if err != nil {
		log.Fatalln(err)
	}
	hashed = pssh.Sum(nil)

	signaturePSS, err := rsa.SignPSS(rand.Reader, privatekey, newhash, hashed, &opts)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("PSS Signature : %x\n", signaturePSS)

	// VerifyPSS
	err = rsa.VerifyPSS(publickey, newhash, hashed, signaturePSS, &opts)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("VerifyPSS successful")
	}

}

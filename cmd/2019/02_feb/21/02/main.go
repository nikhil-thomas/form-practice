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
	params := new(dsa.Parameters)
	if err := dsa.GenerateParameters(
		params,
		rand.Reader,
		dsa.L1024N160,
	); err != nil {
		log.Fatalln(err)
	}

	privateKey := new(dsa.PrivateKey)
	privateKey.PublicKey.Parameters = *params
	dsa.GenerateKey(privateKey, rand.Reader)

	var publicKey dsa.PublicKey
	publicKey = privateKey.PublicKey

	fmt.Println("Private Key")
	fmt.Printf("%x\n", privateKey)

	fmt.Println("public Key")
	fmt.Printf("%x\n", publicKey)

	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	io.WriteString(h, "This message should be signed and then verified")
	signHash := h.Sum(nil)

	r, s, err := dsa.Sign(rand.Reader, privateKey, signHash)
	if err != nil {
		log.Fatalln(err)
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	fmt.Printf("Signature: %x\n", signature)

	verifyStatus := dsa.Verify(&publicKey, signHash, r, s)
	fmt.Println("verifyStatus:", verifyStatus)

	io.WriteString(h, "THis message is not to be signed and verified")
	signHash = h.Sum(nil)
	verifyStatus = dsa.Verify(&publicKey, signHash, r, s)
	fmt.Println("verifyStatus:", verifyStatus)
}

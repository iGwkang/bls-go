package main

import (
	"encoding/hex"
	"fmt"
	"github.com/GwkangZzz/bls-go/bls"
)

func main() {
	// Creating keys and signatures
	// Example seed, used to generate private key. Always use
	// a secure RNG with sufficient entropy to generate a seed (at least 32 bytes).
	seed := []byte{0,  50, 6,  244, 24,  199, 1,  25,  52,  88,  192,
		19, 18, 12, 89,  6,   220, 18, 102, 58,  209, 82,
		12, 62, 89, 110, 182, 9,   44, 20,  254, 22}
	sk := bls.AugSchemeKeyGen(seed)
	pk := sk.GetG1Element()

	message := []byte{1, 2, 3, 4, 5}
	signature:=bls.AugSchemeSign(sk, message)

	// Verify the signature
	ok := bls.AugSchemeVerify(pk, message, signature)
	fmt.Println("Verify the signature is", ok)

	// Serializing keys and signatures to bytes
	skBytes := sk.Serialize()
	pkBytes:=pk.Serialize()
	signatureBytes := signature.Serialize();

	fmt.Println(hex.EncodeToString(skBytes[:]))
	fmt.Println(hex.EncodeToString(pkBytes[:]))
	fmt.Println(hex.EncodeToString(signatureBytes[:]))
}

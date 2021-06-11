package main

import (
	"fmt"
	"github.com/GwkangZzz/bls-go/bls"
)

func main() {
	g1 := bls.NewG1Element()
	fmt.Println(g1.GetFingerprint())
	g1.Serialize()
}

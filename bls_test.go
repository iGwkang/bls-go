package main

import (
	"github.com/GwkangZzz/bls-go/bls"
	"testing"
)

func TestEle(t *testing.T) {
	g1 := bls.NewG1Element()
	t.Log(g1.GetFingerprint())
	g1.Serialize()
}

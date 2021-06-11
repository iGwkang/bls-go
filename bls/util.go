package bls

// #include "util.h"
import "C"
import "unsafe"

func Hash256(message []byte) (hash [32]byte) {
	C.Hash256((*C.uchar)(unsafe.Pointer(&hash[0])), (*C.uchar)(unsafe.Pointer(&message[0])), C.int(len(message)))
	return
}

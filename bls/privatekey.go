package bls

/********************************************************
 * Description : privatekey wrapper
 * Author      : Gwkang
 * Email       : 975500206@qq.com
 * Version     : 1.0
 * History     :
 * Copyright(C): 2021
 ********************************************************/

/*
#include "cprivatekey.h"
*/
import "C"
import (
	"reflect"
	"runtime"
	"unsafe"
)

const PRIVATE_KEY_SIZE = 32

type PrivateKey struct {
	nocopy
	ptr C.CPrivateKey
}

func autoGCPrivateKey(p *PrivateKey) {
	runtime.SetFinalizer(p, func(pk *PrivateKey) {
		C.FreePrivateKey(pk.ptr)
	})
}

func PrivateKeyFromBytes(bytes [PRIVATE_KEY_SIZE]byte) *PrivateKey {
	pk := &PrivateKey{ptr: C.PrivateKeyFromBytes((*C.uchar)(unsafe.Pointer(&bytes[0])))}
	autoGCPrivateKey(pk)
	return pk
}

func (pk *PrivateKey) Serialize() (ret [PRIVATE_KEY_SIZE]byte) {
	ptr := unsafe.Pointer(C.PrivateKeySerialize(pk.ptr))

	var sl []byte
	sp := (*reflect.SliceHeader)(unsafe.Pointer(&sl))
	sp.Data = uintptr(ptr)
	sp.Len = PRIVATE_KEY_SIZE
	sp.Cap = PRIVATE_KEY_SIZE

	copy(ret[:], sl)

	C.free(ptr)

	return
}

func (pk *PrivateKey) Copy() *PrivateKey {
	newPk := &PrivateKey{ptr: C.PrivateKeyCopy(pk.ptr)}
	runtime.SetFinalizer(newPk, func(pk *PrivateKey) {
		C.FreePrivateKey(pk.ptr)
	})
	return newPk
}

func (pk *PrivateKey) GetG1() *G1Element {
	g1 := &G1Element{ptr: C.PrivateKeyGetG1(pk.ptr)}
	return g1
}

func (pk *PrivateKey) GetG2() *G2Element {
	g2 := &G2Element{ptr: C.PrivateKeyGetG2(pk.ptr)}
	return g2
}

func PrivateKeyAggregate(pks []*PrivateKey) *PrivateKey {
	var keys []C.CPrivateKey
	for i := 0; i < len(pks); i++ {
		keys = append(keys, pks[i].ptr)
	}
	if len(keys) == 0 {
		return nil
	}
	cpk := C.PrivateKeyAggregate(&keys[0], C.int(len(pks)))

	newPk := &PrivateKey{ptr: cpk}
	autoGCPrivateKey(newPk)
	return newPk
}

func (pk *PrivateKey) Equal(other *PrivateKey) bool {
	return bool(C.EqualPrivateKey(pk.ptr, other.ptr))
}

func (pk *PrivateKey) MulG1(g1 *G1Element) *G1Element {
	newG1 := &G1Element{ptr: C.MulG1PrivateKey(pk.ptr, g1.ptr)}
	autoGCG1Element(newG1)
	return newG1
}

func (pk *PrivateKey) MulG2(g2 *G2Element) *G2Element {
	newG2 := &G2Element{ptr: C.MulG2PrivateKey(pk.ptr, g2.ptr)}
	autoGCG2Element(newG2)
	return newG2
}

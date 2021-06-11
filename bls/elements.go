package bls

/********************************************************
 * Description : elements wrapper
 * Author      : Gwkang
 * Email       : 975500206@qq.com
 * Version     : 1.0
 * History     :
 * Copyright(C): 2021
 ********************************************************/

/*
#include "celements.h"
*/
import "C"
import (
	"reflect"
	"runtime"
	"unsafe"
)

type G1Element struct {
	nocopy
	ptr C.CG1Element
}

func autoGCG1Element(g1 *G1Element) {
	runtime.SetFinalizer(g1, func(g *G1Element) {
		C.FreeG1Element(g.ptr)
	})
}

func NewG1Element() *G1Element {
	g1 := &G1Element{
		ptr: C.NewG1Element(),
	}
	autoGCG1Element(g1)
	return g1
}

func (g1 *G1Element) Copy() *G1Element {
	newG1 := &G1Element{
		ptr: C.G1ElementCopy(g1.ptr),
	}
	autoGCG1Element(newG1)
	return newG1
}

func G1ElementFromBytes(data []byte) *G1Element {
	g1 := &G1Element{
		ptr: C.G1ElementFromBytes((*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data))),
	}
	autoGCG1Element(g1)
	return g1
}

func G1ElementFromMessage(msg, dst []byte) *G1Element {
	g1 := &G1Element{
		ptr: C.G1ElementFromMessage((*C.uchar)(unsafe.Pointer(&msg[0])), C.int(len(msg)),
			(*C.uchar)(unsafe.Pointer(&dst[0])), C.int(len(dst))),
	}
	autoGCG1Element(g1)
	return g1
}

func G1ElementGenerator() *G1Element {
	g1 := &G1Element{
		ptr: C.G1ElementGenerator(),
	}
	autoGCG1Element(g1)
	return g1
}

func (g1 *G1Element) Negate() *G1Element {
	newG1 := &G1Element{
		ptr: C.G1ElementCopy(g1.ptr),
	}
	autoGCG1Element(newG1)
	return newG1
}

func (g1 *G1Element) GetFingerprint() uint32 {
	return uint32(C.G1ElementGetFingerprint(g1.ptr))
}

func (g1 *G1Element) Serialize() (data []byte) {
	var d *C.uchar
	var size C.uint
	C.G1ElementSerialize(g1.ptr, &d, &size)

	var sl []byte

	sp := (*reflect.SliceHeader)(unsafe.Pointer(&sl))
	sp.Data = uintptr(unsafe.Pointer(d))
	sp.Len = int(size)
	sp.Cap = sp.Len
	copy(data, sl)
	C.free(unsafe.Pointer(d))
	return
}

func (g1 *G1Element) Equal(g *G1Element) bool {
	return bool(C.EqualG1Element(g1.ptr, g.ptr))
}

func (g1 *G1Element) Add(g *G1Element) *G1Element {
	newG1 := &G1Element{
		ptr: C.AddG1Element(g1.ptr, g.ptr),
	}
	autoGCG1Element(newG1)
	return newG1
}

type G2Element struct {
	nocopy
	ptr C.CG2Element
}

func autoGCG2Element(g2 *G2Element) {
	runtime.SetFinalizer(g2, func(g *G2Element) {
		C.FreeG2Element(g.ptr)
	})
}

func NewG2Element() *G2Element {
	g2 := &G2Element{
		ptr: C.NewG2Element(),
	}
	autoGCG2Element(g2)
	return g2
}

func (g2 *G2Element) Copy() *G2Element {
	newG2 := &G2Element{
		ptr: C.G2ElementCopy(g2.ptr),
	}
	autoGCG2Element(newG2)
	return newG2
}

func G2ElementFromBytes(data []byte) *G2Element {
	g2 := &G2Element{
		ptr: C.G2ElementFromBytes((*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data))),
	}
	autoGCG2Element(g2)
	return g2
}

func G2ElementFromMessage(msg, dst []byte) *G2Element {
	g2 := &G2Element{
		ptr: C.G2ElementFromMessage((*C.uchar)(unsafe.Pointer(&msg[0])), C.int(len(msg)),
			(*C.uchar)(unsafe.Pointer(&dst[0])), C.int(len(dst))),
	}
	autoGCG2Element(g2)
	return g2
}

func G2ElementGenerator() *G2Element {
	g2 := &G2Element{
		ptr: C.G2ElementGenerator(),
	}
	autoGCG2Element(g2)
	return g2
}

func (g2 *G2Element) Negate() *G2Element {
	newG2 := &G2Element{
		ptr: C.G2ElementCopy(g2.ptr),
	}
	autoGCG2Element(newG2)
	return newG2
}

func (g2 *G2Element) Serialize() (data []byte) {
	var d *C.uchar
	var size C.uint
	C.G2ElementSerialize(g2.ptr, &d, &size)

	var sl []byte

	sp := (*reflect.SliceHeader)(unsafe.Pointer(&sl))
	sp.Data = uintptr(unsafe.Pointer(d))
	sp.Len = int(size)
	sp.Cap = sp.Len
	copy(data, sl)
	C.free(unsafe.Pointer(d))
	return
}

func (g2 *G2Element) Equal(g *G2Element) bool {
	return bool(C.EqualG2Element(g2.ptr, g.ptr))
}

func (g2 *G2Element) Add(g *G2Element) *G2Element {
	newG2 := &G2Element{
		ptr: C.AddG2Element(g2.ptr, g.ptr),
	}
	autoGCG2Element(newG2)
	return newG2
}

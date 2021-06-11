package bls

/********************************************************
 * Description : schemes wrapper
 * Author      : Gwkang
 * Email       : 975500206@qq.com
 * Version     : 1.0
 * History     :
 * Copyright(C): 2021
 ********************************************************/

/*
#include "cschemes.h"
*/
import "C"
import "unsafe"

func BasicSchemeSkToG1(key *PrivateKey) *G1Element {
	g1 := &G1Element{ptr: C.BasicSchemeSkToG1(key.ptr)}
	autoGCG1Element(g1)
	return g1
}

func BasicSchemeKeyGen(data []byte) *PrivateKey {
	pk := &PrivateKey{ptr: C.BasicSchemeKeyGen((*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data)))}
	autoGCPrivateKey(pk)
	return pk
}

func BasicSchemeDeriveChildSk(key *PrivateKey, index int) *PrivateKey {
	pk := &PrivateKey{ptr: C.BasicSchemeDeriveChildSk(key.ptr, C.int(index))}
	autoGCPrivateKey(pk)
	return pk
}

func BasicSchemeDeriveChildSkUnhardened(key *PrivateKey, index int) *PrivateKey {
	pk := &PrivateKey{ptr: C.BasicSchemeDeriveChildSkUnhardened(key.ptr, C.int(index))}
	autoGCPrivateKey(pk)
	return pk
}

func BasicSchemeDeriveChildPkUnhardened(g1 *G1Element, index int) *G1Element {
	g := &G1Element{ptr: C.BasicSchemeDeriveChildPkUnhardened(g1.ptr, C.int(index))}
	autoGCG1Element(g)
	return g
}

func BasicSchemeAggregateG1(g1 []*G1Element) *G1Element {
	var cg1 []C.CG1Element
	for i := 0; i < len(cg1); i++ {
		cg1 = append(cg1, g1[i].ptr)
	}
	if len(cg1) == 0 {
		return nil
	}
	newG1 := &G1Element{ptr: C.BasicSchemeAggregateG1(&cg1[0], C.int(len(cg1)))}
	autoGCG1Element(newG1)
	return newG1
}

func BasicSchemeAggregateG2(g2 []*G2Element) *G2Element {
	var cg2 []C.CG2Element
	for i := 0; i < len(cg2); i++ {
		cg2 = append(cg2, g2[i].ptr)
	}
	if len(cg2) == 0 {
		return nil
	}
	newG2 := &G2Element{ptr: C.BasicSchemeAggregateG2(&cg2[0], C.int(len(cg2)))}
	autoGCG2Element(newG2)
	return newG2
}

func BasicSchemeSign(key *PrivateKey, data []byte) *G2Element {
	g2 := &G2Element{ptr: C.BasicSchemeSign(key.ptr, (*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data)))}
	autoGCG2Element(g2)
	return g2
}

func BasicSchemeVerify(g1 *G1Element, data []byte, g2 *G2Element) bool {
	return bool(C.BasicSchemeVerify(g1.ptr, (*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data)), g2.ptr))
}

func BasicSchemeeAggregateVerify(g1 []*G1Element, data [][]byte, g2 *G2Element) bool {
	var cg1 []C.CG1Element
	for i := 0; i < len(g1); i++ {
		cg1 = append(cg1, g1[i].ptr)
	}
	var data_vec []C.struct_CData
	for i := 0; i < len(data); i++ {
		cdata := C.struct_CData{}
		cdata.data = (*C.uchar)(unsafe.Pointer(&data[i][0]))
		cdata.len = C.int(len(data[i]))
		data_vec = append(data_vec, cdata)
	}
	if cg1 == nil || data_vec == nil {
		return false
	}

	return bool(C.AugSchemeeAggregateVerify(&cg1[0], C.int(len(cg1)), &data_vec[0], C.int(len(data_vec)), g2.ptr))
}

func AugSchemeSkToG1(key *PrivateKey) *G1Element {
	g1 := &G1Element{ptr: C.AugSchemeSkToG1(key.ptr)}
	autoGCG1Element(g1)
	return g1
}

func AugSchemeKeyGen(data []byte) *PrivateKey {
	pk := &PrivateKey{ptr: C.AugSchemeKeyGen((*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data)))}
	autoGCPrivateKey(pk)
	return pk
}

func AugSchemeDeriveChildSk(key *PrivateKey, index int) *PrivateKey {
	pk := &PrivateKey{ptr: C.AugSchemeDeriveChildSk(key.ptr, C.int(index))}
	autoGCPrivateKey(pk)
	return pk
}

func AugSchemeDeriveChildSkUnhardened(key *PrivateKey, index int) *PrivateKey {
	pk := &PrivateKey{ptr: C.AugSchemeDeriveChildSkUnhardened(key.ptr, C.int(index))}
	autoGCPrivateKey(pk)
	return pk
}

func AugSchemeDeriveChildPkUnhardened(g1 *G1Element, index int) *G1Element {
	g := &G1Element{ptr: C.AugSchemeDeriveChildPkUnhardened(g1.ptr, C.int(index))}
	autoGCG1Element(g)
	return g
}

func AugSchemeAggregateG1(g1 []*G1Element) *G1Element {
	var cg1 []C.CG1Element
	for i := 0; i < len(cg1); i++ {
		cg1 = append(cg1, g1[i].ptr)
	}
	if len(cg1) == 0 {
		return nil
	}
	newG1 := &G1Element{ptr: C.AugSchemeAggregateG1(&cg1[0], C.int(len(cg1)))}
	autoGCG1Element(newG1)
	return newG1
}

func AugSchemeAggregateG2(g2 []*G2Element) *G2Element {
	var cg2 []C.CG2Element
	for i := 0; i < len(cg2); i++ {
		cg2 = append(cg2, g2[i].ptr)
	}
	if len(cg2) == 0 {
		return nil
	}
	newG2 := &G2Element{ptr: C.AugSchemeAggregateG2(&cg2[0], C.int(len(cg2)))}
	autoGCG2Element(newG2)
	return newG2
}

func AugSchemeSign(key *PrivateKey, data []byte) *G2Element {
	g2 := &G2Element{ptr: C.AugSchemeSign(key.ptr, (*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data)))}
	autoGCG2Element(g2)
	return g2
}

func AugSchemeVerify(g1 *G1Element, data []byte, g2 *G2Element) bool {
	return bool(C.AugSchemeVerify(g1.ptr, (*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data)), g2.ptr))
}

func AugSchemeeAggregateVerify(g1 []*G1Element, data [][]byte, g2 *G2Element) bool {
	var cg1 []C.CG1Element
	for i := 0; i < len(g1); i++ {
		cg1 = append(cg1, g1[i].ptr)
	}
	var data_vec []C.struct_CData
	for i := 0; i < len(data); i++ {
		cdata := C.struct_CData{}
		cdata.data = (*C.uchar)(unsafe.Pointer(&data[i][0]))
		cdata.len = C.int(len(data[i]))
		data_vec = append(data_vec, cdata)
	}
	if cg1 == nil || data_vec == nil {
		return false
	}

	return bool(C.AugSchemeeAggregateVerify(&cg1[0], C.int(len(cg1)), &data_vec[0], C.int(len(data_vec)), g2.ptr))
}

func PopSchemeSkToG1(key *PrivateKey) *G1Element {
	g1 := &G1Element{ptr: C.PopSchemeSkToG1(key.ptr)}
	autoGCG1Element(g1)
	return g1
}

func PopSchemeKeyGen(data []byte) *PrivateKey {
	pk := &PrivateKey{ptr: C.PopSchemeKeyGen((*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data)))}
	autoGCPrivateKey(pk)
	return pk
}

func PopSchemeDeriveChildSk(key *PrivateKey, index int) *PrivateKey {
	pk := &PrivateKey{ptr: C.PopSchemeDeriveChildSk(key.ptr, C.int(index))}
	autoGCPrivateKey(pk)
	return pk
}

func PopSchemeDeriveChildSkUnhardened(key *PrivateKey, index int) *PrivateKey {
	pk := &PrivateKey{ptr: C.PopSchemeDeriveChildSkUnhardened(key.ptr, C.int(index))}
	autoGCPrivateKey(pk)
	return pk
}

func PopSchemeDeriveChildPkUnhardened(g1 *G1Element, index int) *G1Element {
	g := &G1Element{ptr: C.PopSchemeDeriveChildPkUnhardened(g1.ptr, C.int(index))}
	autoGCG1Element(g)
	return g
}

func PopSchemeAggregateG1(g1 []*G1Element) *G1Element {
	var cg1 []C.CG1Element
	for i := 0; i < len(cg1); i++ {
		cg1 = append(cg1, g1[i].ptr)
	}
	if len(cg1) == 0 {
		return nil
	}
	newG1 := &G1Element{ptr: C.PopSchemeAggregateG1(&cg1[0], C.int(len(cg1)))}
	autoGCG1Element(newG1)
	return newG1
}

func PopSchemeAggregateG2(g2 []*G2Element) *G2Element {
	var cg2 []C.CG2Element
	for i := 0; i < len(cg2); i++ {
		cg2 = append(cg2, g2[i].ptr)
	}
	if len(cg2) == 0 {
		return nil
	}
	newG2 := &G2Element{ptr: C.PopSchemeAggregateG2(&cg2[0], C.int(len(cg2)))}
	autoGCG2Element(newG2)
	return newG2
}

func PopSchemeSign(key *PrivateKey, data []byte) *G2Element {
	g2 := &G2Element{ptr: C.PopSchemeSign(key.ptr, (*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data)))}
	autoGCG2Element(g2)
	return g2
}

func PopSchemeVerify(g1 *G1Element, data []byte, g2 *G2Element) bool {
	return bool(C.PopSchemeVerify(g1.ptr, (*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data)), g2.ptr))
}

func PopSchemeeAggregateVerify(g1 []*G1Element, data [][]byte, g2 *G2Element) bool {
	var cg1 []C.CG1Element
	for i := 0; i < len(g1); i++ {
		cg1 = append(cg1, g1[i].ptr)
	}
	var data_vec []C.struct_CData
	for i := 0; i < len(data); i++ {
		cdata := C.struct_CData{}
		cdata.data = (*C.uchar)(unsafe.Pointer(&data[i][0]))
		cdata.len = C.int(len(data[i]))
		data_vec = append(data_vec, cdata)
	}
	if cg1 == nil || data_vec == nil {
		return false
	}

	return bool(C.PopSchemeeAggregateVerify(&cg1[0], C.int(len(cg1)), &data_vec[0], C.int(len(data_vec)), g2.ptr))
}

func PopSchemePopProve(key *PrivateKey) *G2Element {
	newG2 := &G2Element{ptr: C.PopSchemePopProve(key.ptr)}
	autoGCG2Element(newG2)
	return newG2
}
func PopSchemePopVerify(g1 *G1Element, g2 *G2Element) bool {
	return bool(C.PopSchemePopVerify(g1.ptr, g2.ptr))
}
func PopSchemeFastAggregateVerify(g1 []*G1Element, data []byte, g2 *G2Element) bool {
	var cg1 []C.CG1Element
	for i := 0; i < len(g1); i++ {
		cg1 = append(cg1, g1[i].ptr)
	}
	if cg1 == nil {
		return false
	}

	return bool(C.PopSchemeFastAggregateVerify(&cg1[0], C.int(len(cg1)), (*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data)), g2.ptr))
}

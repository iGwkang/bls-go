package bls

/********************************************************
 * Description : bls wrapper
 * Author      : Gwkang
 * Email       : 975500206@qq.com
 * Version     : 1.0
 * History     :
 * Copyright(C): 2021
 ********************************************************/

/*
#cgo CFLAGS: -I./relic/include -std=c11
#cgo CXXFLAGS: -I./relic/include -std=c++17
#cgo LDFLAGS: -L${SRCDIR}/relic/lib -lrelic_s -lstdc++
//#cgo windows LDFLAGS: -l${SRCDIR}/relic/lib/relic_s -lstdc++
#include <stdlib.h>
*/
import "C"

type nocopy struct{}

func (*nocopy) Lock() {}
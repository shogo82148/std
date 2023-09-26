// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

type I1 interface {
	Method1()
}

type I2 interface {
	Method1()
	Method2()
}

type TS uint16
type TM uintptr
type TL [2]uintptr

type T8 uint8
type T16 uint16
type T32 uint32
type T64 uint64
type Tstr string
type Tslice []byte

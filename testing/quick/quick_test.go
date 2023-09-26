// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package quick

type TestBoolAlias bool

type TestFloat32Alias float32

type TestFloat64Alias float64

type TestComplex64Alias complex64

type TestComplex128Alias complex128

type TestInt16Alias int16

type TestInt32Alias int32

type TestInt64Alias int64

type TestInt8Alias int8

type TestIntAlias int

type TestMapAlias map[int]int

type TestSliceAlias []byte

type TestStringAlias string

type TestStruct struct {
	A int
	B string
}

type TestStructAlias TestStruct

type TestUint16Alias uint16

type TestUint32Alias uint32

type TestUint64Alias uint64

type TestUint8Alias uint8

type TestUintAlias uint

type TestUintptrAlias uintptr

type TestIntptrAlias *int

// This tests that ArbitraryValue is working by checking that all the arbitrary
// values of type MyStruct have x = 42.

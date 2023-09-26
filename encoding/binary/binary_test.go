// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binary

type Struct struct {
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128
	Array      [4]uint8
	Bool       bool
	BoolArray  [4]bool
}

type T struct {
	Int     int
	Uint    uint
	Uintptr uintptr
	Array   [4]int
}

// Addresses of arrays are easier to manipulate with reflection than are slices.

type BlankFields struct {
	A uint32
	_ int32
	B float64
	_ [4]int16
	C byte
	_ [7]byte
	_ struct {
		f [8]float32
	}
}

type BlankFieldsProbe struct {
	A  uint32
	P0 int32
	B  float64
	P1 [4]int16
	C  byte
	P2 [7]byte
	P3 struct {
		F [8]float32
	}
}

type Unexported struct {
	a int32
}

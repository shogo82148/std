// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.regabiargs

package reflect_test

// As of early May 2021 this is no longer necessary for amd64,
// but it remains in case this is needed for the next register abi port.
// TODO (1.18) If enabling register ABI on additional architectures turns out not to need this, remove it.
type MagicLastTypeNameForTestingRegisterABI struct{}

type StructWithMethods struct {
	Value int
}

type StructFewRegs struct {
	a0, a1, a2, a3 int
	f0, f1, f2, f3 float64
}

type StructFillRegs struct {
	a0, a1, a2, a3, a4, a5, a6, a7, a8                              int
	f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14 float64
}

// Struct1 is a simple integer-only aggregate struct.
type Struct1 struct {
	A, B, C uint
}

// Struct2 is Struct1 but with an array-typed field that will
// force it to get passed on the stack.
type Struct2 struct {
	A, B, C uint
	D       [2]uint32
}

// Struct3 is Struct2 but with an anonymous array-typed field.
// This should act identically to Struct2.
type Struct3 struct {
	A, B, C uint
	D       [2]uint32
}

// Struct4 has byte-length fields that should
// each use up a whole registers.
type Struct4 struct {
	A, B int8
	C, D uint8
	E    bool
}

// Struct5 is a relatively large struct
// with both integer and floating point values.
type Struct5 struct {
	A             uint16
	B             int16
	C, D          uint32
	E             int32
	F, G, H, I, J float32
}

// Struct6 has a nested struct.
type Struct6 struct {
	Struct1
}

// Struct7 is a struct with a nested array-typed field
// that cannot be passed in registers as a result.
type Struct7 struct {
	Struct1
	Struct2
}

// Struct8 is large aggregate struct type that may be
// passed in registers.
type Struct8 struct {
	Struct5
	Struct1
}

// Struct9 is a type that has an array type nested
// 2 layers deep, and as a result needs to be passed
// on the stack.
type Struct9 struct {
	Struct1
	Struct7
}

// Struct10 is a struct type that is too large to be
// passed in registers.
type Struct10 struct {
	Struct5
	Struct8
}

// Struct11 is a struct type that has several reference
// types in it.
type Struct11 struct {
	X map[string]int
}

// Struct12 has Struct11 embedded into it to test more
// paths.
type Struct12 struct {
	A int
	Struct11
}

// Struct13 tests an empty field.
type Struct13 struct {
	A int
	X struct{}
	B int
}

// Struct14 tests a non-zero-sized (and otherwise register-assignable)
// struct with a field that is a non-zero length array with zero-sized members.
type Struct14 struct {
	A uintptr
	X [3]struct{}
	B float64
}

// Struct15 tests a non-zero-sized (and otherwise register-assignable)
// struct with a struct field that is zero-sized but contains a
// non-zero length array with zero-sized members.
type Struct15 struct {
	A uintptr
	X struct {
		Y [3]struct{}
	}
	B float64
}

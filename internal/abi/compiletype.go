// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abi

// CommonSize returns sizeof(Type) for a compilation target with a given ptrSize
func CommonSize(ptrSize int) int

// StructFieldSize returns sizeof(StructField) for a compilation target with a given ptrSize
func StructFieldSize(ptrSize int) int

// UncommonSize returns sizeof(UncommonType).  This currently does not depend on ptrSize.
// This exported function is in an internal package, so it may change to depend on ptrSize in the future.
func UncommonSize() uint64

// IMethodSize returns sizeof(IMethod) for a compilation target with a given ptrSize
func IMethodSize(ptrSize int) int

// TFlagOff returns the offset of Type.TFlag for a compilation target with a given ptrSize
func TFlagOff(ptrSize int) int

// Offset is for computing offsets of type data structures at compile/link time;
// the target platform may not be the host platform.  Its state includes the
// current offset, necessary alignment for the sequence of types, and the size
// of pointers and alignment of slices, interfaces, and strings (this is for tearing-
// resistant access to these types, if/when that is supported).
type Offset struct {
	off        uint64
	align      uint8
	ptrSize    uint8
	sliceAlign uint8
}

// NewOffset returns a new Offset with offset 0 and alignment 1.
func NewOffset(ptrSize uint8, twoWordAlignSlices bool) Offset

// InitializedOffset returns a new Offset with specified offset, alignment, pointer size, and slice alignment.
func InitializedOffset(off int, align uint8, ptrSize uint8, twoWordAlignSlices bool) Offset

// Align returns the offset obtained by aligning offset to a multiple of a.
// a must be a power of two.
func (o Offset) Align(a uint8) Offset

// D8 returns the offset obtained by appending an 8-bit field to o.
func (o Offset) D8() Offset

// D16 returns the offset obtained by appending a 16-bit field to o.
func (o Offset) D16() Offset

// D32 returns the offset obtained by appending a 32-bit field to o.
func (o Offset) D32() Offset

// D64 returns the offset obtained by appending a 64-bit field to o.
func (o Offset) D64() Offset

// D64 returns the offset obtained by appending a pointer field to o.
func (o Offset) P() Offset

// Slice returns the offset obtained by appending a slice field to o.
func (o Offset) Slice() Offset

// String returns the offset obtained by appending a string field to o.
func (o Offset) String() Offset

// Interface returns the offset obtained by appending an interface field to o.
func (o Offset) Interface() Offset

// Offset returns the struct-aligned offset (size) of o.
// This is at least as large as the current internal offset; it may be larger.
func (o Offset) Offset() uint64

func (o Offset) PlusUncommon() Offset

// CommonOffset returns the Offset to the data after the common portion of type data structures.
func CommonOffset(ptrSize int, twoWordAlignSlices bool) Offset

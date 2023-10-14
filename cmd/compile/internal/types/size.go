// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

var PtrSize int

var RegSize int

// Slices in the runtime are represented by three components:
//
//	type slice struct {
//		ptr unsafe.Pointer
//		len int
//		cap int
//	}
//
// Strings in the runtime are represented by two components:
//
//	type string struct {
//		ptr unsafe.Pointer
//		len int
//	}
//
// These variables are the offsets of fields and sizes of these structs.
var (
	SlicePtrOffset int64
	SliceLenOffset int64
	SliceCapOffset int64

	SliceSize  int64
	StringSize int64
)

var SkipSizeForTracing bool

// MaxWidth is the maximum size of a value on the target architecture.
var MaxWidth int64

// CalcSizeDisabled indicates whether it is safe
// to calculate Types' widths and alignments. See CalcSize.
var CalcSizeDisabled bool

// RoundUp rounds o to a multiple of r, r is a power of 2.
func RoundUp(o int64, r int64) int64

// CalcSize calculates and stores the size and alignment for t.
// If CalcSizeDisabled is set, and the size/alignment
// have not already been calculated, it calls Fatal.
// This is used to prevent data races in the back end.
func CalcSize(t *Type)

// CalcStructSize calculates the size of t,
// filling in t.width, t.align, t.intRegs, and t.floatRegs,
// even if size calculation is otherwise disabled.
func CalcStructSize(t *Type)

func CheckSize(t *Type)

func DeferCheckSize()

func ResumeCheckSize()

// PtrDataSize returns the length in bytes of the prefix of t
// containing pointer data. Anything after this offset is scalar data.
//
// PtrDataSize is only defined for actual Go types. It's an error to
// use it on compiler-internal types (e.g., TSSA, TRESULTS).
func PtrDataSize(t *Type) int64

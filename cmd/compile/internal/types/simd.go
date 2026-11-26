// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// IsSIMD reports whether t is a SIMD type: one of the SSA vector and mask
// types, such as TypeVec128, or one of the vector and mask types of
// simd/archsimd, or a type defined in terms of one of those, such as
// bridge.Int32x4 in package simd/internal/bridge, which is defined as
// archsimd.Int32x4. Values of SIMD types live in vector registers. The tag
// types that mark the vector and mask types of simd/archsimd are not SIMD
// types.
func (t *Type) IsSIMD() bool

// ContainsSIMD reports whether t is a SIMD type, or an array or struct
// that contains one.
func (t *Type) ContainsSIMD() bool

// SIMDShape describes the elements of a SIMD vector or mask type.
type SIMDShape struct {
	// Elem is the element type. For example, it's float32 for Float32x4.
	//
	// The Elem of a mask is the signed integer type of the size of the
	// vector elements it governs: int32 for Mask32x4 and for Mask32s.
	Elem *Type

	// Lanes is the number of elements of a fixed-width type, or 0 for a
	// scalable type.
	Lanes int64

	// Scalable reports whether this is an arm64 SVE type, whose number of
	// elements is only known at run time.
	Scalable bool

	// Mask reports whether this is a mask type, such as Mask32x4 or
	// Mask32s.
	//
	// The fixed-width masks are stored as vectors whose elements are 0 or
	// -1, so most of the compiler can treat them like vectors. But a
	// scalable mask on arm64 is an SVE predicate: it's stored as a bitmap
	// with one bit per byte of a scalable vector rather than as an array
	// of Elem, and lives in a predicate register.
	Mask bool
}

// SIMDShape returns the shape of SIMD vector or mask type t, and whether t is
// such a type. It works on the vector and mask types of simd/archsimd, and
// types defined in terms of them, such as the types of package simd on
// platforms where it isn't emulated. Each is a struct whose first field is a
// SIMD tag and whose second field holds its elements.
func (t *Type) SIMDShape() (shape SIMDShape, ok bool)

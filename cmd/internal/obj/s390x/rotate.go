// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s390x

// RotateParams represents the immediates required for a "rotate
// then ... selected bits instruction".
//
// The Start and End values are the indexes that represent
// the masked region. They are inclusive and are in big-
// endian order (bit 0 is the MSB, bit 63 is the LSB). They
// may wrap around.
//
// Some examples:
//
// Masked region             | Start | End
// --------------------------+-------+----
// 0x00_00_00_00_00_00_00_0f | 60    | 63
// 0xf0_00_00_00_00_00_00_00 | 0     | 3
// 0xf0_00_00_00_00_00_00_0f | 60    | 3
//
// The Amount value represents the amount to rotate the
// input left by. Note that this rotation is performed
// before the masked region is used.
type RotateParams struct {
	Start  uint8
	End    uint8
	Amount uint8
}

// NewRotateParams creates a set of parameters representing a
// rotation left by the amount provided and a selection of the bits
// between the provided start and end indexes (inclusive).
//
// The start and end indexes and the rotation amount must all
// be in the range 0-63 inclusive or this function will panic.
func NewRotateParams(start, end, amount uint8) RotateParams

// RotateLeft generates a new set of parameters with the rotation amount
// increased by the given value. The selected bits are left unchanged.
func (r RotateParams) RotateLeft(amount uint8) RotateParams

// OutMask provides a mask representing the selected bits.
func (r RotateParams) OutMask() uint64

// InMask provides a mask representing the selected bits relative
// to the source value (i.e. pre-rotation).
func (r RotateParams) InMask() uint64

// OutMerge tries to generate a new set of parameters representing
// the intersection between the selected bits and the provided mask.
// If the intersection is unrepresentable (0 or not contiguous) nil
// will be returned.
func (r RotateParams) OutMerge(mask uint64) *RotateParams

// InMerge tries to generate a new set of parameters representing
// the intersection between the selected bits and the provided mask
// as applied to the source value (i.e. pre-rotation).
// If the intersection is unrepresentable (0 or not contiguous) nil
// will be returned.
func (r RotateParams) InMerge(mask uint64) *RotateParams

func (RotateParams) CanBeAnSSAAux()

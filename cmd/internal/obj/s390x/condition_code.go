// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s390x

// CCMask represents a 4-bit condition code mask. Bits that
// are not part of the mask should be 0.
//
// Condition code masks represent the 4 possible values of
// the 2-bit condition code as individual bits. Since IBM Z
// is a big-endian platform bits are numbered from left to
// right. The lowest value, 0, is represented by 8 (0b1000)
// and the highest value, 3, is represented by 1 (0b0001).
//
// Note that condition code values have different semantics
// depending on the instruction that set the condition code.
// The names given here assume that the condition code was
// set by an integer or floating point comparison. Other
// instructions may use these same codes to indicate
// different results such as a carry or overflow.
type CCMask uint8

const (
	Never CCMask = 0

	// 1-bit masks
	Equal     CCMask = 1 << 3
	Less      CCMask = 1 << 2
	Greater   CCMask = 1 << 1
	Unordered CCMask = 1 << 0

	// 2-bit masks
	EqualOrUnordered   CCMask = Equal | Unordered
	LessOrEqual        CCMask = Less | Equal
	LessOrGreater      CCMask = Less | Greater
	LessOrUnordered    CCMask = Less | Unordered
	GreaterOrEqual     CCMask = Greater | Equal
	GreaterOrUnordered CCMask = Greater | Unordered

	// 3-bit masks
	NotEqual     CCMask = Always ^ Equal
	NotLess      CCMask = Always ^ Less
	NotGreater   CCMask = Always ^ Greater
	NotUnordered CCMask = Always ^ Unordered

	// 4-bit mask
	Always CCMask = Equal | Less | Greater | Unordered

	// useful aliases
	Carry    CCMask = GreaterOrUnordered
	NoCarry  CCMask = LessOrEqual
	Borrow   CCMask = NoCarry
	NoBorrow CCMask = Carry
)

// Inverse returns the complement of the condition code mask.
func (c CCMask) Inverse() CCMask

// ReverseComparison swaps the bits at 0b0100 and 0b0010 in the mask,
// reversing the behavior of greater than and less than conditions.
func (c CCMask) ReverseComparison() CCMask

func (c CCMask) String() string

func (CCMask) CanBeAnSSAAux()

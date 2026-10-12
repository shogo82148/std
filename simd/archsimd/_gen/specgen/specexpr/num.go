// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpr

// Num represents a number in the solver. This is abstracted because we work
// with both concrete numbers and *symbolic widths* and need to be able to mix
// them. A Num is also an [Expr].
type Num interface {
	Expr

	ValidWidth() bool

	Mul(x Num) (Num, error)
	Div(x Num) (Num, error)
	Compare(o Num) (int, bool)

	String() string
}

// Int is an integer that satisfies [Num].
type Int int

func (w Int) ValidWidth() bool

func (w Int) Mul(x Num) (Num, error)

func (w Int) Div(x Num) (Num, error)

func (w Int) Compare(o Num) (int, bool)

func (w Int) String() string

// An ScalableWidth represents a symbolic width relative to a fixed but unknown
// scalable vector width VW. This is represented as a rational factor
// VW*num/denom
type ScalableWidth struct {
	num, denom int
}

// VW returns the base ScalableWidth representing a full-width scalable vector.
func VW() ScalableWidth

func (w ScalableWidth) ValidWidth() bool

func (w ScalableWidth) Mul(x Num) (Num, error)

func (w ScalableWidth) Div(x Num) (Num, error)

func (w ScalableWidth) Compare(o Num) (int, bool)

func (w ScalableWidth) String() string

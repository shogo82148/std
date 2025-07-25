// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/cmd/internal/src"
)

// DivisionNeedsFixUp reports whether the division needs fix-up code.
func DivisionNeedsFixUp(v *Value) bool

// Aux is an interface to hold miscellaneous data in Blocks and Values.
type Aux interface {
	CanBeAnSSAAux()
}

var AuxMark auxMark

func StringToAux(s string) Aux

func IsInlinableMemmove(dst, src *Value, sz int64, c *Config) bool

func LogLargeCopy(funcName string, pos src.XPos, s int64)

func GetPPC64Shiftsh(auxint int64) int64

func GetPPC64Shiftmb(auxint int64) int64

func GetPPC64Shiftme(auxint int64) int64

// DecodePPC64RotateMask is the inverse operation of encodePPC64RotateMask.  The values returned as
// mb and me satisfy the POWER ISA definition of MASK(x,y) where MASK(mb,me) = mask.
func DecodePPC64RotateMask(sauxint int64) (rotate, mb, me int64, mask uint64)

// PanicBoundsC contains a constant for a bounds failure.
type PanicBoundsC struct {
	C int64
}

// PanicBoundsCC contains 2 constants for a bounds failure.
type PanicBoundsCC struct {
	Cx int64
	Cy int64
}

func (p PanicBoundsC) CanBeAnSSAAux()

func (p PanicBoundsCC) CanBeAnSSAAux()

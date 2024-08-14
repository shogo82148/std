// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loong64

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

const (
	FuncAlign = 4
)

type Optab struct {
	as    obj.As
	from1 uint8
	reg   uint8
	from3 uint8
	to1   uint8
	to2   uint8
	type_ int8
	size  int8
	param int16
	flag  uint8
}

const (
	NOTUSETMP = 1 << iota
)

func IsAtomicInst(as obj.As) bool

// r1 -> rk
// r2 -> rj
// r3 -> rd
func OP_RRR(op uint32, r1 uint32, r2 uint32, r3 uint32) uint32

// r2 -> rj
// r3 -> rd
func OP_RR(op uint32, r2 uint32, r3 uint32) uint32

func OP_16IR_5I(op uint32, i uint32, r2 uint32) uint32

func OP_16IRR(op uint32, i uint32, r2 uint32, r3 uint32) uint32

func OP_12IRR(op uint32, i uint32, r2 uint32, r3 uint32) uint32

func OP_IR(op uint32, i uint32, r2 uint32) uint32

func OP_15I(op uint32, i uint32) uint32

// Encoding for the 'b' or 'bl' instruction.
func OP_B_BL(op uint32, i uint32) uint32

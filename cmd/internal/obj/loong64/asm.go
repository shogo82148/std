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

// The constants here define the data characteristics within the bit field range.
//
//	ALL1: The data in the bit field is all 1
//	ALL0: The data in the bit field is all 0
//	ST1: The data in the bit field starts with 1, but not all 1
//	ST0: The data in the bit field starts with 0, but not all 0
const (
	ALL1 = iota
	ALL0
	ST1
	ST0
)

func OP_RRRR(op uint32, r1 uint32, r2 uint32, r3 uint32, r4 uint32) uint32

// r1 -> rk
// r2 -> rj
// r3 -> rd
func OP_RRR(op uint32, r1 uint32, r2 uint32, r3 uint32) uint32

// r2 -> rj
// r3 -> rd
func OP_RR(op uint32, r2 uint32, r3 uint32) uint32

func OP_16IR_5I(op uint32, i uint32, r2 uint32) uint32

func OP_16IRR(op uint32, i uint32, r2 uint32, r3 uint32) uint32

func OP_12IR_5I(op uint32, i1 uint32, r2 uint32, i2 uint32) uint32

func OP_12IRR(op uint32, i uint32, r2 uint32, r3 uint32) uint32

func OP_8IRR(op uint32, i uint32, r2 uint32, r3 uint32) uint32

func OP_6IRR(op uint32, i uint32, r2 uint32, r3 uint32) uint32

func OP_5IRR(op uint32, i uint32, r2 uint32, r3 uint32) uint32

func OP_4IRR(op uint32, i uint32, r2 uint32, r3 uint32) uint32

func OP_3IRR(op uint32, i uint32, r2 uint32, r3 uint32) uint32

func OP_IR(op uint32, i uint32, r2 uint32) uint32

func OP_15I(op uint32, i uint32) uint32

// i1 -> msb
// r2 -> rj
// i3 -> lsb
// r4 -> rd
func OP_IRIR(op uint32, i1 uint32, r2 uint32, i3 uint32, r4 uint32) uint32

// Encoding for the 'b' or 'bl' instruction.
func OP_B_BL(op uint32, i uint32) uint32

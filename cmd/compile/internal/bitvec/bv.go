// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitvec

import (
	"github.com/shogo82148/std/cmd/internal/src"
)

// A BitVec is a bit vector.
type BitVec struct {
	N int32
	B []uint32
}

func New(n int32) BitVec

type Bulk struct {
	words []uint32
	nbit  int32
	nword int32
}

func NewBulk(nbit int32, count int32, pos src.XPos) Bulk

func (b *Bulk) Next() BitVec

func (bv1 BitVec) Eq(bv2 BitVec) bool

func (dst BitVec) Copy(src BitVec)

func (bv BitVec) Get(i int32) bool

func (bv BitVec) Set(i int32)

func (bv BitVec) Unset(i int32)

// bvnext returns the smallest index >= i for which bvget(bv, i) == 1.
// If there is no such index, bvnext returns -1.
func (bv BitVec) Next(i int32) int32

func (bv BitVec) IsEmpty() bool

func (bv BitVec) Count() int

func (bv BitVec) Not()

// union
func (dst BitVec) Or(src1, src2 BitVec)

// intersection
func (dst BitVec) And(src1, src2 BitVec)

// difference
func (dst BitVec) AndNot(src1, src2 BitVec)

func (bv BitVec) String() string

func (bv BitVec) Clear()

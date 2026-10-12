// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssaop

// A RegMask encodes a set of machine registers.
type RegMask struct {
	V1, V2 uint64
}

type Register uint8

func (r RegMask) Intersect(s RegMask) RegMask

func (r RegMask) Union(s RegMask) RegMask

func (r RegMask) Minus(s RegMask) RegMask

func (r RegMask) Empty() bool

func (r RegMask) PickReg() Register

func (r RegMask) AddReg(i Register) RegMask

func (r RegMask) RemoveReg(i Register) RegMask

func (r RegMask) HasReg(i Register) bool

func (m RegMask) String() string

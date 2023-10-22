// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objw

import (
	"github.com/shogo82148/std/cmd/compile/internal/bitvec"
	"github.com/shogo82148/std/cmd/internal/obj"
)

// Uint8 writes an unsigned byte v into s at offset off,
// and returns the next unused offset (i.e., off+1).
func Uint8(s *obj.LSym, off int, v uint8) int

func Uint16(s *obj.LSym, off int, v uint16) int

func Uint32(s *obj.LSym, off int, v uint32) int

func Uintptr(s *obj.LSym, off int, v uint64) int

// Uvarint writes a varint v into s at offset off,
// and returns the next unused offset.
func Uvarint(s *obj.LSym, off int, v uint64) int

func Bool(s *obj.LSym, off int, v bool) int

// UintN writes an unsigned integer v of size wid bytes into s at offset off,
// and returns the next unused offset.
func UintN(s *obj.LSym, off int, v uint64, wid int) int

func SymPtr(s *obj.LSym, off int, x *obj.LSym, xoff int) int

func SymPtrWeak(s *obj.LSym, off int, x *obj.LSym, xoff int) int

func SymPtrOff(s *obj.LSym, off int, x *obj.LSym) int

func SymPtrWeakOff(s *obj.LSym, off int, x *obj.LSym) int

func Global(s *obj.LSym, width int32, flags int16)

// BitVec writes the contents of bv into s as sequence of bytes
// in little-endian order, and returns the next unused offset.
func BitVec(s *obj.LSym, off int, bv bitvec.BitVec) int

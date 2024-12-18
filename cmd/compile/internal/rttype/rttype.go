// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rttype allows the compiler to share type information with
// the runtime. The shared type information is stored in
// internal/abi. This package translates those types from the host
// machine on which the compiler runs to the target machine on which
// the compiled program will run. In particular, this package handles
// layout differences between e.g. a 64 bit compiler and 32 bit
// target.
package rttype

import (
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
)

// The type structures shared with the runtime.
var Type *types.Type

var ArrayType *types.Type
var ChanType *types.Type
var FuncType *types.Type
var InterfaceType *types.Type
var OldMapType *types.Type
var SwissMapType *types.Type
var PtrType *types.Type
var SliceType *types.Type
var StructType *types.Type

// Types that are parts of the types above.
var IMethod *types.Type
var Method *types.Type
var StructField *types.Type
var UncommonType *types.Type

// Type switches and asserts
var InterfaceSwitch *types.Type
var TypeAssert *types.Type

// Interface tables (itabs)
var ITab *types.Type

func Init()

// A Cursor represents a typed location inside a static variable where we
// are going to write.
type Cursor struct {
	lsym   *obj.LSym
	offset int64
	typ    *types.Type
}

// NewCursor returns a cursor starting at lsym+off and having type t.
func NewCursor(lsym *obj.LSym, off int64, t *types.Type) Cursor

// WritePtr writes a pointer "target" to the component at the location specified by c.
func (c Cursor) WritePtr(target *obj.LSym)

func (c Cursor) WritePtrWeak(target *obj.LSym)

func (c Cursor) WriteUintptr(val uint64)

func (c Cursor) WriteUint32(val uint32)

func (c Cursor) WriteUint16(val uint16)

func (c Cursor) WriteUint8(val uint8)

func (c Cursor) WriteInt(val int64)

func (c Cursor) WriteInt32(val int32)

func (c Cursor) WriteBool(val bool)

// WriteSymPtrOff writes a "pointer" to the given symbol. The symbol
// is encoded as a uint32 offset from the start of the section.
func (c Cursor) WriteSymPtrOff(target *obj.LSym, weak bool)

// WriteSlice writes a slice header to c. The pointer is target+off, the len and cap fields are given.
func (c Cursor) WriteSlice(target *obj.LSym, off, len, cap int64)

// Reloc adds a relocation from the current cursor position.
// Reloc fills in Off and Siz fields. Caller should fill in the rest (Type, others).
func (c Cursor) Reloc(rel obj.Reloc)

// Field selects the field with the given name from the struct pointed to by c.
func (c Cursor) Field(name string) Cursor

func (c Cursor) Elem(i int64) Cursor

type ArrayCursor struct {
	c Cursor
	n int
}

// NewArrayCursor returns a cursor starting at lsym+off and having n copies of type t.
func NewArrayCursor(lsym *obj.LSym, off int64, t *types.Type, n int) ArrayCursor

// Elem selects element i of the array pointed to by c.
func (a ArrayCursor) Elem(i int) Cursor

// ModifyArray converts a cursor pointing at a type [k]T to a cursor pointing
// at a type [n]T.
// Also returns the size delta, aka (n-k)*sizeof(T).
func (c Cursor) ModifyArray(n int) (ArrayCursor, int64)

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goobj

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/internal/abi"
)

// CUFileIndex is used to index the filenames that are stored in the
// per-package/per-CU FileList.
type CUFileIndex uint32

// FuncInfo is serialized as a symbol (aux symbol). The symbol data is
// the binary encoding of the struct below.
type FuncInfo struct {
	Args      uint32
	Locals    uint32
	FuncID    abi.FuncID
	FuncFlag  abi.FuncFlag
	StartLine int32
	File      []CUFileIndex
	InlTree   []InlTreeNode
}

func (a *FuncInfo) Write(w *bytes.Buffer)

// FuncInfoLengths is a cache containing a roadmap of offsets and
// lengths for things within a serialized FuncInfo. Each length field
// stores the number of items (e.g. files, inltree nodes, etc), and the
// corresponding "off" field stores the byte offset of the start of
// the items in question.
type FuncInfoLengths struct {
	NumFile     uint32
	FileOff     uint32
	NumInlTree  uint32
	InlTreeOff  uint32
	Initialized bool
}

func (*FuncInfo) ReadFuncInfoLengths(b []byte) FuncInfoLengths

func (*FuncInfo) ReadArgs(b []byte) uint32

func (*FuncInfo) ReadLocals(b []byte) uint32

func (*FuncInfo) ReadFuncID(b []byte) abi.FuncID

func (*FuncInfo) ReadFuncFlag(b []byte) abi.FuncFlag

func (*FuncInfo) ReadStartLine(b []byte) int32

func (*FuncInfo) ReadFile(b []byte, filesoff uint32, k uint32) CUFileIndex

func (*FuncInfo) ReadInlTree(b []byte, inltreeoff uint32, k uint32) InlTreeNode

// InlTreeNode is the serialized form of FileInfo.InlTree.
type InlTreeNode struct {
	Parent   int32
	File     CUFileIndex
	Line     int32
	Func     SymRef
	ParentPC int32
}

func (inl *InlTreeNode) Write(w *bytes.Buffer)

// Read an InlTreeNode from b, return the remaining bytes.
func (inl *InlTreeNode) Read(b []byte) []byte

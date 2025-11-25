// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssagen

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
)

// SymABIs records information provided by the assembler about symbol
// definition ABIs and reference ABIs.
type SymABIs struct {
	defs map[string]obj.ABI
	refs map[string]obj.ABISet
}

func NewSymABIs() *SymABIs

// ReadSymABIs reads a symabis file that specifies definitions and
// references of text symbols by ABI.
//
// The symabis format is a set of lines, where each line is a sequence
// of whitespace-separated fields. The first field is a verb and is
// either "def" for defining a symbol ABI or "ref" for referencing a
// symbol using an ABI. For both "def" and "ref", the second field is
// the symbol name and the third field is the ABI name, as one of the
// named cmd/internal/obj.ABI constants.
func (s *SymABIs) ReadSymABIs(file string)

// HasDef returns whether the given symbol has an assembly definition.
func (s *SymABIs) HasDef(sym *types.Sym) bool

// GenABIWrappers applies ABI information to Funcs and generates ABI
// wrapper functions where necessary.
func (s *SymABIs) GenABIWrappers()

// CreateWasmImportWrapper creates a wrapper for imported WASM functions to
// adapt them to the Go calling convention. The body for this function is
// generated in cmd/internal/obj/wasm/wasmobj.go
func CreateWasmImportWrapper(fn *ir.Func) bool

func GenWasmExportWrapper(wrapped *ir.Func)

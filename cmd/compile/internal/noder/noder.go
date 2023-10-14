// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package noder

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
	"github.com/shogo82148/std/cmd/compile/internal/types"
)

func LoadPackage(filenames []string)

// WasmImport stores metadata associated with the //go:wasmimport pragma
type WasmImport struct {
	Pos    syntax.Pos
	Module string
	Name   string
}

func Renameinit() *types.Sym

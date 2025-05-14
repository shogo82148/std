// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file encapsulates some of the odd characteristics of the
// Loong64 (LoongArch64) instruction set, to minimize its interaction
// with the core of the assembler.

package arch

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

// IsLoong64MUL reports whether the op (as defined by an loong64.A* constant) is
// one of the MUL/DIV/REM instructions that require special handling.
func IsLoong64MUL(op obj.As) bool

// IsLoong64RDTIME reports whether the op (as defined by an loong64.A*
// constant) is one of the RDTIMELW/RDTIMEHW/RDTIMED instructions that
// require special handling.
func IsLoong64RDTIME(op obj.As) bool

func IsLoong64PRELD(op obj.As) bool

func IsLoong64AMO(op obj.As) bool

// Loong64RegisterExtension constructs an Loong64 register with extension or arrangement.
func Loong64RegisterExtension(a *obj.Addr, ext string, reg, num int16, isAmount, isIndex bool) error

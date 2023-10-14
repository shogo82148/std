// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file encapsulates some of the odd characteristics of the
// 64-bit PowerPC (PPC64) instruction set, to minimize its interaction
// with the core of the assembler.

package arch

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

// IsPPC64CMP reports whether the op (as defined by an ppc64.A* constant) is
// one of the CMP instructions that require special handling.
func IsPPC64CMP(op obj.As) bool

// IsPPC64NEG reports whether the op (as defined by an ppc64.A* constant) is
// one of the NEG-like instructions that require special handling.
func IsPPC64NEG(op obj.As) bool

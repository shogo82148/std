// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssacompile

import (
	"github.com/shogo82148/std/cmd/compile/internal/ssa"
)

// IsGlobalAddr reports whether v is known to be an address of a global (or nil).
func IsGlobalAddr(v *ssa.Value) bool

// IsReadOnlyGlobalAddr reports whether v is known to be an address of a read-only global.
func IsReadOnlyGlobalAddr(v *ssa.Value) bool

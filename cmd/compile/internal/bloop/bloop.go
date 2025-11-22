// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bloop

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

// BloopWalk performs a walk on all functions in the package
// if it imports testing and wrap the results of all qualified
// statements in a runtime.KeepAlive intrinsic call. See package
// doc for more details.
//
//	for b.Loop() {...}
//
// loop's body.
func BloopWalk(pkg *ir.Package)

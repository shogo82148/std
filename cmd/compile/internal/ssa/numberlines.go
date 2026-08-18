// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import "github.com/shogo82148/std/cmd/compile/internal/ssa/ssaop"

// NotStmtBoundary reports whether a value with opcode op can never be a statement
// boundary. Such values don't correspond to a user's understanding of a
// statement boundary.
func NotStmtBoundary(op ssaop.Op) bool

func (b *Block) FirstPossibleStmtValue() *Value

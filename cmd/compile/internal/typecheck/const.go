// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typecheck

import (
	"github.com/shogo82148/std/go/constant"

	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
)

func DefaultLit(n ir.Node, t *types.Type) ir.Node

// OrigConst returns an OLITERAL with orig n and value v.
func OrigConst(n ir.Node, v constant.Value) ir.Node

func OrigBool(n ir.Node, v bool) ir.Node

func OrigInt(n ir.Node, v int64) ir.Node

// IndexConst checks if Node n contains a constant expression
// representable as a non-negative int and returns its value.
// If n is not a constant expression, not representable as an
// integer, or negative, it returns -1. If n is too large, it
// returns -2.
func IndexConst(n ir.Node) int64

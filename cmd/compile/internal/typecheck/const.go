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

// ConvertVal converts v into a representation appropriate for t. If
// no such representation exists, it returns constant.MakeUnknown()
// instead.
//
// If explicit is true, then conversions from integer to string are
// also allowed.
func ConvertVal(v constant.Value, t *types.Type, explicit bool) constant.Value

// IndexConst checks if Node n contains a constant expression
// representable as a non-negative int and returns its value.
// If n is not a constant expression, not representable as an
// integer, or negative, it returns -1. If n is too large, it
// returns -2.
func IndexConst(n ir.Node) int64

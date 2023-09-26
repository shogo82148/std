// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements printing of expressions.

package types

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/go/ast"
)

// ExprString returns the (possibly simplified) string representation for x.
func ExprString(x ast.Expr) string

// WriteExpr writes the (possibly simplified) string representation for x to buf.
func WriteExpr(buf *bytes.Buffer, x ast.Expr)

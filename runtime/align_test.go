// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

type Visitor struct {
	fset    *token.FileSet
	types   map[ast.Expr]types.TypeAndValue
	checked map[string]bool
	t       *testing.T
}

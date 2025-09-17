// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is a copy of golang.org/x/tools/internal/astutil/clone.go

package astutil

import (
	"github.com/shogo82148/std/go/ast"
)

// CloneNode returns a deep copy of a Node.
// It omits pointers to ast.{Scope,Object} variables.
func CloneNode[T ast.Node](n T) T

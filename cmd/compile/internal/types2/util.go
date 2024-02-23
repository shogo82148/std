// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains various functionality that is
// different between go/types and types2. Factoring
// out this code allows more of the rest of the code
// to be shared.

package types2

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
)

// ExprString returns a string representation of x.
func ExprString(x syntax.Node) string

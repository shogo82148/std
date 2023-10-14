// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package escape

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

// Fmt is called from node printing to print information about escape analysis results.
func Fmt(n ir.Node) string

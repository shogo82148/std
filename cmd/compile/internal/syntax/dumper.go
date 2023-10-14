// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements printing of syntax tree structures.

package syntax

import (
	"github.com/shogo82148/std/io"
)

// Fdump dumps the structure of the syntax tree rooted at n to w.
// It is intended for debugging purposes; no specific output format
// is guaranteed.
func Fdump(w io.Writer, n Node) (err error)

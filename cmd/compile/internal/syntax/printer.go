// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements printing of syntax trees in source format.

package syntax

import (
	"github.com/shogo82148/std/io"
)

// Form controls print formatting.
type Form uint

const (
	_ Form = iota
	LineForm
	ShortForm
)

// Fprint prints node x to w in the specified form.
// It returns the number of bytes written, and whether there was an error.
func Fprint(w io.Writer, x Node, form Form) (n int, err error)

// String is a convenience function that prints n in ShortForm
// and returns the printed string.
func String(n Node) string

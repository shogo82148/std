// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux || darwin

package cryptotest

import (
	"github.com/shogo82148/std/testing"
)

// BoundarySlices allocates a pair of slices of the given size one at the start
// of a page, another at the end. Any access beyond or before the slice
// boundaries should cause a fault.
func BoundarySlices(t *testing.T, size int) (start, end []byte)

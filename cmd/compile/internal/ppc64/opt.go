// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ppc64

// Many Power ISA arithmetic and logical instructions come in four
// standard variants. These bits let us map between variants.
const (
	V_CC = 1 << 0
	V_V  = 1 << 1
)

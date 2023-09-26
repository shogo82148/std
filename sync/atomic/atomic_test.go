// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic_test

import (
	. "sync/atomic"
)

// Do the 64-bit functions panic? If so, don't bother testing.

// Test that this compiles.
// When atomic.Pointer used _ [0]T, it did not.
type List struct {
	Next Pointer[List]
}

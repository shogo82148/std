// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug_test

import (
	. "runtime/debug"
)

type Obj struct {
	x, y int
}

type G[T any] struct{}
type I interface {
	M()
}

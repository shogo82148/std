// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// pprofのようなプロファイルの提供。

package main

import (
	"github.com/shogo82148/std/internal/trace"
)

// Record represents one entry in pprof-like profiles.
type Record struct {
	stk  []*trace.Frame
	n    uint64
	time int64
}

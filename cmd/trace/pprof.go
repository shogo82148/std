// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// pprofのようなプロファイルの提供。

package main

// Recordはpprofのようなプロファイルのエントリを表します。
type Record struct {
	stk  []*trace.Frame
	n    uint64
	time int64
}

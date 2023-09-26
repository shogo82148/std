// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic_test

import (
	. "sync/atomic"
)

var Value_SwapTests = []struct {
	init interface{}
	new  interface{}
	want interface{}
	err  interface{}
}{
	{init: nil, new: nil, err: "sync/atomic: swap of nil value into Value"},
	{init: nil, new: true, want: nil, err: nil},
	{init: true, new: "", err: "sync/atomic: swap of inconsistently typed value into Value"},
	{init: true, new: false, want: true, err: nil},
}

var Value_CompareAndSwapTests = []struct {
	init interface{}
	new  interface{}
	old  interface{}
	want bool
	err  interface{}
}{
	{init: nil, new: nil, old: nil, err: "sync/atomic: compare and swap of nil value into Value"},
	{init: nil, new: true, old: "", err: "sync/atomic: compare and swap of inconsistently typed values into Value"},
	{init: nil, new: true, old: true, want: false, err: nil},
	{init: nil, new: true, old: nil, want: true, err: nil},
	{init: true, new: "", err: "sync/atomic: compare and swap of inconsistently typed value into Value"},
	{init: true, new: true, old: false, want: false, err: nil},
	{init: true, new: true, old: true, want: true, err: nil},
	{init: heapA, new: struct{ uint }{1}, old: heapB, want: true, err: nil},
}

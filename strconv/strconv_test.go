// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv_test

// Sink makes sure the compiler cannot optimize away the benchmarks.
var Sink struct {
	Bool       bool
	Int        int
	Int64      int64
	Uint64     uint64
	Float64    float64
	Complex128 complex128
	Error      error
	Bytes      []byte
}

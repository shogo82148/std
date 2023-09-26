// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package quick

type TestStruct struct {
	A int
	B string
}

// This tests that ArbitraryValue is working by checking that all the arbitrary
// values of type MyStruct have x = 42.

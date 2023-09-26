// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanner

// A StringReader delivers its data one string segment at a time via Read.
type StringReader struct {
	data []string
	step int
}

// An errReader returns (0, err) where err is not io.EOF.

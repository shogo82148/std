// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unicode_test

import (
	. "unicode"
)

type T struct {
	rune   rune
	script string
}

// Hand-chosen tests from Unicode 5.1.0, 6.0.0, 6.2.0, 6.3.0, 7.0.0, 8.0.0,
// 9.0.0, 10.0.0.
// mostly to discover when new scripts and categories arise.
// If this tests fails, add the missing scripts to the test and add entries
// of the form
//     pkg unicode, var <new script> *RangeTable
// to api/next.txt.

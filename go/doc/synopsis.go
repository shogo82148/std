// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doc

// Synopsis returns a cleaned version of the first sentence in s.
// That sentence ends after the first period followed by space and
// not preceded by exactly one uppercase letter. The result string
// has no \n, \r, or \t characters and uses only single spaces between
// words.
func Synopsis(s string) string

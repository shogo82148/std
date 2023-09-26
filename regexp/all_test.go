// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package regexp

type ReplaceTest struct {
	pattern, replacement, input, output string
}

type ReplaceFuncTest struct {
	pattern       string
	replacement   func(string) string
	input, output string
}

type MetaTest struct {
	pattern, output, literal string
	isLiteral                bool
}

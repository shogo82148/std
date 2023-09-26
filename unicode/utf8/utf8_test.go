// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utf8_test

import (
	. "unicode/utf8"
)

type Utf8Map struct {
	r   rune
	str string
}

type RuneCountTest struct {
	in  string
	out int
}

type RuneLenTest struct {
	r    rune
	size int
}

type ValidTest struct {
	in  string
	out bool
}

type ValidRuneTest struct {
	r  rune
	ok bool
}

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package url

type URLTest struct {
	in        string
	out       *URL
	roundtrip string
}

type EscapeTest struct {
	in  string
	out string
	err error
}

type EncodeQueryTest struct {
	m        Values
	expected string
}

type RequestURITest struct {
	url *URL
	out string
}

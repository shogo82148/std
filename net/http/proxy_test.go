// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

var UseProxyTests = []struct {
	host  string
	match bool
}{

	{"localhost:80", false},
	{"127.0.0.1", false},
	{"127.0.0.2", false},
	{"[::1]", false},
	{"[::2]", true},

	{"barbaz.net", false},
	{"foobar.com", false},
	{"foofoobar.com", true},
	{"baz.com", true},
	{"localhost.net", true},
	{"local.localhost", true},
	{"barbarbaz.net", true},
	{"www.foobar.com", true},
}

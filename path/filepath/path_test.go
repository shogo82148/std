// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filepath_test

type PathTest struct {
	path, result string
}

type SplitListTest struct {
	list   string
	result []string
}

type SplitTest struct {
	path, dir, file string
}

type JoinTest struct {
	elem []string
	path string
}

type ExtTest struct {
	path, ext string
}

type Node struct {
	name    string
	entries []*Node
	mark    int
}

type IsAbsTest struct {
	path  string
	isAbs bool
}

type EvalSymlinksTest struct {
	path, dest string
}

var EvalSymlinksTestDirs = []EvalSymlinksTest{
	{"test", ""},
	{"test/dir", ""},
	{"test/dir/link3", "../../"},
	{"test/link1", "../test"},
	{"test/link2", "dir"},
	{"test/linkabs", "/"},
}

var EvalSymlinksTests = []EvalSymlinksTest{
	{"test", "test"},
	{"test/dir", "test/dir"},
	{"test/dir/../..", "."},
	{"test/link1", "test"},
	{"test/link2", "test/dir"},
	{"test/link1/dir", "test/dir"},
	{"test/link2/..", "test"},
	{"test/dir/link3", "."},
	{"test/link2/link3/test", "test"},
	{"test/linkabs", "/"},
}

var EvalSymlinksAbsWindowsTests = []EvalSymlinksTest{
	{`c:\`, `c:\`},
}

// Test directories relative to temporary directory.
// The tests are run in absTestDirs[0].

// Test paths relative to temporary directory. $ expands to the directory.
// The tests are run in absTestDirs[0].
// We create absTestDirs first.

type RelTests struct {
	root, path, want string
}

type VolumeNameTest struct {
	path string
	vol  string
}

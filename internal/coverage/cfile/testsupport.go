// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cfile

import (
	"github.com/shogo82148/std/io"
)

// ProcessCoverTestDir is called from
// testmain code when "go test -cover" is in effect. It is not
// intended to be used other than internally by the Go command's
// generated code.
func ProcessCoverTestDir(dir string, cfile string, cm string, cpkg string, w io.Writer, selpkgs []string) error

// Snapshot returns a snapshot of coverage percentage at a moment of
// time within a running test, so as to support the testing.Coverage()
// function. This version doesn't examine coverage meta-data, so the
// result it returns will be less accurate (more "slop") due to the
// fact that we don't look at the meta data to see how many statements
// are associated with each counter.
func Snapshot() float64

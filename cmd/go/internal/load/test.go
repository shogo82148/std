// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package load

import (
	"github.com/shogo82148/std/context"
)

var TestMainDeps = []string{

	"os",
	"reflect",
	"testing",
	"testing/internal/testdeps",
}

type TestCover struct {
	Mode  string
	Local bool
	Pkgs  []*Package
	Paths []string
	Vars  []coverInfo
}

// TestPackagesFor is like TestPackagesAndErrors but it returns
// an error if the test packages or their dependencies have errors.
// Only test packages without errors are returned.
func TestPackagesFor(ctx context.Context, opts PackageOpts, p *Package, cover *TestCover) (pmain, ptest, pxtest *Package, err error)

// TestPackagesAndErrors returns three packages:
//   - pmain, the package main corresponding to the test binary (running tests in ptest and pxtest).
//   - ptest, the package p compiled with added "package p" test files.
//   - pxtest, the result of compiling any "package p_test" (external) test files.
//
// If the package has no "package p_test" test files, pxtest will be nil.
// If the non-test compilation of package p can be reused
// (for example, if there are no "package p" test files and
// package p need not be instrumented for coverage or any other reason),
// then the returned ptest == p.
//
// If done is non-nil, TestPackagesAndErrors will finish filling out the returned
// package structs in a goroutine and call done once finished. The members of the
// returned packages should not be accessed until done is called.
//
// The caller is expected to have checked that len(p.TestGoFiles)+len(p.XTestGoFiles) > 0,
// or else there's no point in any of this.
func TestPackagesAndErrors(ctx context.Context, done func(), opts PackageOpts, p *Package, cover *TestCover) (pmain, ptest, pxtest *Package)

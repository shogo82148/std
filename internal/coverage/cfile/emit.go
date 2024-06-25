// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cfile implements management of coverage files.
// It provides functionality exported in runtime/coverage as well as
// additional functionality used directly by package testing
// through testing/internal/testdeps.
package cfile

// MarkProfileEmitted signals the coverage machinery that
// coverage data output files have already been written out, and there
// is no need to take any additional action at exit time. This
// function is called from the coverage-related boilerplate code in _testmain.go
// emitted for go unit tests.
func MarkProfileEmitted(val bool)

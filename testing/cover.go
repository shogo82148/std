// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Support for test coverage.

package testing

// CoverBlock records the coverage data for a single basic block.
// NOTE: This struct is internal to the testing infrastructure and may change.
// It is not covered (yet) by the Go 1 compatibility guidelines.
type CoverBlock struct {
	Line0 uint32
	Col0  uint16
	Line1 uint32
	Col1  uint16
	Stmts uint16
}

// Cover records information about test coverage checking.
// NOTE: This struct is internal to the testing infrastructure and may change.
// It is not covered (yet) by the Go 1 compatibility guidelines.
type Cover struct {
	Mode            string
	Counters        map[string][]uint32
	Blocks          map[string][]CoverBlock
	CoveredPackages string
}

// Coverage reports the current code coverage as a fraction in the range [0, 1].
// If coverage is not enabled, Coverage returns 0.
//
// When running a large set of sequential test cases, checking Coverage after each one
// can be useful for identifying which test cases exercise new code paths.
// It is not a replacement for the reports generated by 'go test -cover' and
// 'go tool cover'.
func Coverage() float64

// RegisterCover records the coverage data accumulators for the tests.
// NOTE: This function is internal to the testing infrastructure and may change.
// It is not covered (yet) by the Go 1 compatibility guidelines.
func RegisterCover(c Cover)

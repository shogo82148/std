// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fipsdit provides an analyzer which checks that exported functions in
// crypto/... do not call into crypto/internal/fips140/... or constant-time
// scalar primitives (crypto/subtle.ConstantTime{Select,ByteEq,Eq,LessOrEq} and
// crypto/internal/constanttime) without enabling Data Independent Timing via
// crypto/subtle.WithDataIndependentTiming.
package fipsdit

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:      "fipsditcheck",
	Doc:       "check FIPS 140 Data Independent Timing (DIT) requirements",
	Requires:  []*analysis.Analyzer{inspect.Analyzer},
	FactTypes: []analysis.Fact{(*funcFact)(nil)},
	Run:       run,
}

// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cformat

import (
	"github.com/shogo82148/std/internal/coverage"
	"github.com/shogo82148/std/io"
)

type Formatter struct {
	// Maps import path to package state.
	pm map[string]*pstate
	// Records current package being visited.
	pkg string
	// Pointer to current package state.
	p *pstate
	// Counter mode.
	cm coverage.CounterMode
}

func NewFormatter(cm coverage.CounterMode) *Formatter

// SetPackage tells the formatter that we're about to visit the
// coverage data for the package with the specified import path.
// Note that it's OK to call SetPackage more than once with the
// same import path; counter data values will be accumulated.
func (fm *Formatter) SetPackage(importpath string)

// AddUnit passes info on a single coverable unit (file, funcname,
// literal flag, range of lines, and counter value) to the formatter.
// Counter values will be accumulated where appropriate.
func (fm *Formatter) AddUnit(file string, fname string, isfnlit bool, unit coverage.CoverableUnit, count uint32)

// EmitTextual writes the accumulated coverage data in the legacy
// cmd/cover text format to the writer 'w'. We sort the data items by
// importpath, source file, and line number before emitting (this sorting
// is not explicitly mandated by the format, but seems like a good idea
// for repeatable/deterministic dumps).
func (fm *Formatter) EmitTextual(w io.Writer) error

// EmitPercent writes out a "percentage covered" string to the writer
// 'w', selecting the set of packages in 'pkgs' and suffixing the
// printed string with 'inpkgs'.
func (fm *Formatter) EmitPercent(w io.Writer, pkgs []string, inpkgs string, noteEmpty bool, aggregate bool) error

// EmitFuncs writes out a function-level summary to the writer 'w'. A
// note on handling function literals: although we collect coverage
// data for unnamed literals, it probably does not make sense to
// include them in the function summary since there isn't any good way
// to name them (this is also consistent with the legacy cmd/cover
// implementation). We do want to include their counts in the overall
// summary however.
func (fm *Formatter) EmitFuncs(w io.Writer) error

// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/io"
)

// Compile is the main entry point for this package.
// Compile modifies f so that on return:
//   - all Values in f map to 0 or 1 assembly instructions of the target architecture
//   - the order of f.Blocks is the order to emit the Blocks
//   - the order of b.Values is the order to emit the Values in each Block
//   - f has a non-nil regAlloc field
func Compile(f *Func)

// DumpFileForPhase creates a file from the function name and phase name,
// warning and returning nil if this is not possible.
func (f *Func) DumpFileForPhase(phaseName string) io.WriteCloser

// Debug output
var IntrinsicsDebug int
var IntrinsicsDisable bool

var BuildDebug int
var BuildTest int
var BuildStats int
var BuildDump map[string]bool = make(map[string]bool)

var GenssaDump map[string]bool = make(map[string]bool)

// PhaseOption sets the specified flag in the specified ssa phase,
// returning empty string if this was successful or a string explaining
// the error if it was not.
// A version of the phase name with "_" replaced by " " is also checked for a match.
// If the phase name begins a '~' then the rest of the underscores-replaced-with-blanks
// version is used as a regular expression to match the phase name(s).
//
// Special cases that have turned out to be useful:
//   - ssa/check/on enables checking after each phase
//   - ssa/all/time enables time reporting for all phases
//
// See gc/lex.go for dissection of the option string.
// Example uses:
//
// GO_GCFLAGS=-d=ssa/generic_cse/time,ssa/generic_cse/stats,ssa/generic_cse/debug=3 ./make.bash
//
// BOOT_GO_GCFLAGS=-d='ssa/~^.*scc$/off' GO_GCFLAGS='-d=ssa/~^.*scc$/off' ./make.bash
func PhaseOption(phase, flag string, val int, valString string) string

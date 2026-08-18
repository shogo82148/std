// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssacompile

import (
	"github.com/shogo82148/std/cmd/compile/internal/ssa"
)

// Compiler satisfies the ssacore.Compiler interface.
type Compiler struct{}

func (_ Compiler) Passes() []ssa.Pass

// Compile is the main entry point for this package.
// Compile modifies f so that on return:
//   - all Values in f map to 0 or 1 assembly instructions of the target architecture
//   - the order of f.Blocks is the order to emit the Blocks
//   - the order of b.Values is the order to emit the Values in each Block
//   - f has a non-nil regAlloc field
func (_ Compiler) Compile(f *ssa.Func, htmlWriter ssa.HTMLWriter)

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

func PostCompile()

// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/cmd/internal/src"
	"github.com/shogo82148/std/internal/bisect"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
)

type HashDebug struct {
	mu   sync.Mutex
	name string
	// what file (if any) receives the yes/no logging?
	// default is os.Stdout
	logfile          io.Writer
	posTmp           []src.Pos
	bytesTmp         bytes.Buffer
	matches          []hashAndMask
	excludes         []hashAndMask
	bisect           *bisect.Matcher
	fileSuffixOnly   bool
	inlineSuffixOnly bool
}

// SetInlineSuffixOnly controls whether hashing and reporting use the entire
// inline position, or just the most-inline suffix.  Compiler debugging tends
// to want the whole inlining, debugging user problems (loopvarhash, e.g.)
// typically does not need to see the entire inline tree, there is just one
// copy of the source code.
func (d *HashDebug) SetInlineSuffixOnly(b bool) *HashDebug

var ConvertHash *HashDebug
var FmaHash *HashDebug
var LoopVarHash *HashDebug
var PGOHash *HashDebug
var LiteralAllocHash *HashDebug
var MergeLocalsHash *HashDebug
var VariableMakeHash *HashDebug

// DebugHashMatchPkgFunc reports whether debug variable Gossahash
//
//  1. is empty (returns true; this is a special more-quickly implemented case of 4 below)
//
//  2. is "y" or "Y" (returns true)
//
//  3. is "n" or "N" (returns false)
//
//  4. does not explicitly exclude the sha1 hash of pkgAndName (see step 6)
//
//  5. is a suffix of the sha1 hash of pkgAndName (returns true)
//
//  6. OR
//     if the (non-empty) value is in the regular language
//     "(-[01]+/)+?([01]+(/[01]+)+?"
//     (exclude..)(....include...)
//     test the [01]+ exclude substrings, if any suffix-match, return false (4 above)
//     test the [01]+ include substrings, if any suffix-match, return true
//     The include substrings AFTER the first slash are numbered 0,1, etc and
//     are named fmt.Sprintf("%s%d", varname, number)
//     As an extra-special case for multiple failure search,
//     an excludes-only string ending in a slash (terminated, not separated)
//     implicitly specifies the include string "0/1", that is, match everything.
//     (Exclude strings are used for automated search for multiple failures.)
//     Clause 6 is not really intended for human use and only
//     matters for failures that require multiple triggers.
//
// Otherwise it returns false.
//
// Unless Flags.Gossahash is empty, when DebugHashMatchPkgFunc returns true the message
//
//	"%s triggered %s\n", varname, pkgAndName
//
// is printed on the file named in environment variable GSHS_LOGFILE,
// or standard out if that is empty.  "Varname" is either the name of
// the variable or the name of the substring, depending on which matched.
//
// Typical use:
//
//  1. you make a change to the compiler, say, adding a new phase
//
//  2. it is broken in some mystifying way, for example, make.bash builds a broken
//     compiler that almost works, but crashes compiling a test in run.bash.
//
//  3. add this guard to the code, which by default leaves it broken, but does not
//     run the broken new code if Flags.Gossahash is non-empty and non-matching:
//
//     if !base.DebugHashMatch(ir.PkgFuncName(fn)) {
//     return nil // early exit, do nothing
//     }
//
//  4. rebuild w/o the bad code,
//     GOCOMPILEDEBUG=gossahash=n ./all.bash
//     to verify that you put the guard in the right place with the right sense of the test.
//
//  5. use github.com/dr2chase/gossahash to search for the error:
//
//     go install github.com/dr2chase/gossahash@latest
//
//     gossahash -- <the thing that fails>
//
//     for example: GOMAXPROCS=1 gossahash -- ./all.bash
//
//  6. gossahash should return a single function whose miscompilation
//     causes the problem, and you can focus on that.
func DebugHashMatchPkgFunc(pkg, fn string) bool

func DebugHashMatchPos(pos src.XPos) bool

// HasDebugHash returns true if Flags.Gossahash is non-empty, which
// results in hashDebug being not-nil.  I.e., if !HasDebugHash(),
// there is no need to create the string for hashing and testing.
func HasDebugHash() bool

// NewHashDebug returns a new hash-debug tester for the
// environment variable ev.  If ev is not set, it returns
// nil, allowing a lightweight check for normal-case behavior.
func NewHashDebug(ev, s string, file io.Writer) *HashDebug

// MatchPkgFunc returns true if either the variable used to create d is
// unset, or if its value is y, or if it is a suffix of the base-two
// representation of the hash of pkg and fn.  If the variable is not nil,
// then a true result is accompanied by stylized output to d.logfile, which
// is used for automated bug search.
func (d *HashDebug) MatchPkgFunc(pkg, fn string, note func() string) bool

// MatchPos is similar to MatchPkgFunc, but for hash computation
// it uses the source position including all inlining information instead of
// package name and path.
// Note that the default answer for no environment variable (d == nil)
// is "yes", do the thing.
func (d *HashDebug) MatchPos(pos src.XPos, desc func() string) bool

// MatchPosWithInfo is similar to MatchPos, but with additional information
// that is included for hash computation, so it can distinguish multiple
// matches on the same source location.
// Note that the default answer for no environment variable (d == nil)
// is "yes", do the thing.
func (d *HashDebug) MatchPosWithInfo(pos src.XPos, info any, desc func() string) bool

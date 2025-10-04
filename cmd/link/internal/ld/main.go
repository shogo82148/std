// Inferno utils/6l/obj.c
// https://bitbucket.org/inferno-os/inferno-os/src/master/utils/6l/obj.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ld

import (
	"github.com/shogo82148/std/cmd/internal/sys"
	"github.com/shogo82148/std/flag"
)

// Flags used by the linker. The exported flags are used by the architecture-specific packages.
var (
	FlagC = flag.Bool("c", false, "dump call graph")
	FlagD = flag.Bool("d", false, "disable dynamic executable")

	FlagS = flag.Bool("s", false, "disable symbol table")

	FlagDebugTramp    = flag.Int("debugtramp", 0, "debug trampolines")
	FlagDebugTextSize = flag.Int("debugtextsize", 0, "debug text section max size")

	FlagStrictDups = flag.Int("strictdups", 0, "sanity check duplicate symbol contents during object file reading (1=warn 2=err).")
	FlagRound      = flag.Int64("R", -1, "set address rounding `quantum`")
	FlagTextAddr   = flag.Int64("T", -1, "set the start address of text symbols")
	FlagFuncAlign  = flag.Int("funcalign", 0, "set function align to `N` bytes")

	FlagW = new(bool)
)

// Main is the main entry point for the linker code.
func Main(arch *sys.Arch, theArch Arch)

type Rpath struct {
	set bool
	val string
}

func (r *Rpath) Set(val string) error

func (r *Rpath) String() string

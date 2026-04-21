// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package flags implements top-level flags and the usage message for the assembler.
package flags

import (
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/flag"
)

var (
	Debug      = flag.Bool("debug", false, "dump instructions as they are parsed")
	OutputFile = flag.String("o", "", "output file; default foo.o for /a/b/c/foo.s as first argument")
	TrimPath   = flag.String("trimpath", "", "remove prefix from recorded source file paths")
	Shared     = flag.Bool("shared", false, "generate code that can be linked into a shared library")
	Dynlink    = flag.Bool("dynlink", false, "support references to Go symbols defined in other shared libraries")
	Linkshared = flag.Bool("linkshared", false, "generate code that will be linked against Go shared libraries")
	AllErrors  = flag.Bool("e", false, "no limit on number of errors reported")
	SymABIs    = flag.Bool("gensymabis", false, "write symbol ABI information to output file, don't assemble")
	Importpath = flag.String("p", obj.UnlinkablePkg, "set expected package import to path")
	Spectre    = flag.String("spectre", "", "enable spectre mitigations in `list` (all, ret)")
)

var DebugFlags struct {
	CompressInstructions int    `help:"use compressed instructions when possible (if supported by architecture)"`
	MayMoreStack         string `help:"call named function before all stack growth checks"`
	PCTab                string `help:"print named pc-value table\nOne of: pctospadj, pctofile, pctoline, pctoinline, pctopcdata"`
}

var (
	D        MultiFlag
	I        MultiFlag
	PrintOut int
	DebugV   bool
)

// MultiFlag allows setting a value multiple times to collect a list, as in -I=dir1 -I=dir2.
type MultiFlag []string

func (m *MultiFlag) String() string

func (m *MultiFlag) Set(val string) error

func Usage()

func Parse()

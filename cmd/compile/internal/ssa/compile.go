// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"
)

type Compiler interface {
	Compile(f *Func, htmlWriter HTMLWriter)
	Passes() []Pass
}

type HTMLWriter interface {
	Enabled() bool
	FlushPhases()
	WritePhase(phase, title string)
	WriteColumn(phase, title, class, html string)
	DebugInfo(v func(*Value) string)
	TimeFormatting() time.Duration
	Close()
}

type Pass struct {
	Name     string
	Fn       func(*Func)
	Required bool
	Disabled bool
	Time     bool
	Mem      bool
	Stats    int
	Debug    int
	Test     int
	Dump     map[string]bool
	Keywords map[string]int64
	UsedKW   map[string]bool
}

// DumpFileForPhase creates a file from the function name and phase name,
// warning and returning nil if this is not possible.
func (f *Func) DumpFileForPhase(phaseName string) io.WriteCloser

// DumpFile creates a file from the phase name and function name
// Dumping is done to files to avoid buffering huge strings before
// output.
func (f *Func) DumpFile(phaseName string)

func (p *Pass) AddDump(s string)

func (p *Pass) String() string

func (p *Pass) Val(kw string, ifUnset int64) int64

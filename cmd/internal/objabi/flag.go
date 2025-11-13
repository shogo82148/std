// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objabi

import (
	"github.com/shogo82148/std/io"
)

func Flagcount(name, usage string, val *int)

func Flagfn1(name, usage string, f func(string))

func Flagprint(w io.Writer)

func Flagparse(usage func())

func AddVersionFlag()

// DecodeArg decodes an argument.
//
// This function is public for testing with the parallel encoder.
func DecodeArg(arg string) string

type DebugFlag struct {
	tab          map[string]debugField
	concurrentOk *bool
	debugSSA     DebugSSA
}

// A DebugSSA function is called to set a -d ssa/... option.
// If nil, those options are reported as invalid options.
// If DebugSSA returns a non-empty string, that text is reported as a compiler error.
// If phase is "help", it should print usage information and terminate the process.
type DebugSSA func(phase, flag string, val int, valString string) string

// NewDebugFlag constructs a DebugFlag for the fields of debug, which
// must be a pointer to a struct.
//
// Each field of *debug is a different value, named for the lower-case of the field name.
// Each field must be an int or string and must have a `help` struct tag.
// There may be an "Any bool" field, which will be set if any debug flags are set.
//
// The returned flag takes a comma-separated list of settings.
// Each setting is name=value; for ints, name is short for name=1.
//
// If debugSSA is non-nil, any debug flags of the form ssa/... will be
// passed to debugSSA for processing.
func NewDebugFlag(debug any, debugSSA DebugSSA) *DebugFlag

func (f *DebugFlag) Set(debugstr string) error

func (f *DebugFlag) String() string

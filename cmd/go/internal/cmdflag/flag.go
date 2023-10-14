// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cmdflag handles flag processing common to several go tools.
package cmdflag

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/flag"
)

// ErrFlagTerminator indicates the distinguished token "--", which causes the
// flag package to treat all subsequent arguments as non-flags.
var ErrFlagTerminator = errors.New("flag terminator")

// A FlagNotDefinedError indicates a flag-like argument that does not correspond
// to any registered flag in a FlagSet.
type FlagNotDefinedError struct {
	RawArg   string
	Name     string
	HasValue bool
	Value    string
}

func (e FlagNotDefinedError) Error() string

// A NonFlagError indicates an argument that is not a syntactically-valid flag.
type NonFlagError struct {
	RawArg string
}

func (e NonFlagError) Error() string

// ParseOne sees if args[0] is present in the given flag set and if so,
// sets its value and returns the flag along with the remaining (unused) arguments.
//
// ParseOne always returns either a non-nil Flag or a non-nil error,
// and always consumes at least one argument (even on error).
//
// Unlike (*flag.FlagSet).Parse, ParseOne does not log its own errors.
func ParseOne(fs *flag.FlagSet, args []string) (f *flag.Flag, remainingArgs []string, err error)

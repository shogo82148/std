// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

import (
	"github.com/shogo82148/std/flag"
)

// GOFLAGS returns the flags from $GOFLAGS.
// The list can be assumed to contain one string per flag,
// with each string either beginning with -name or --name.
func GOFLAGS() []string

// InitGOFLAGS initializes the goflags list from $GOFLAGS.
// If goflags is already initialized, it does nothing.
func InitGOFLAGS()

// SetFromGOFLAGS sets the flags in the given flag set using settings in $GOFLAGS.
func SetFromGOFLAGS(flags *flag.FlagSet, ignoreErrors bool)

// InGOFLAGS returns whether GOFLAGS contains the given flag, such as "-mod".
func InGOFLAGS(flag string) bool

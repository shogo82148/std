// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !cmd_go_bootstrap && !compiler_bootstrap

package counter

import (
	"github.com/shogo82148/std/flag"

	"golang.org/x/telemetry/counter"
)

func OpenCalled() bool

// Open opens the counter files for writing if telemetry is supported
// on the current platform (and does nothing otherwise).
func Open()

// Inc increments the counter with the given name.
func Inc(name string)

// New returns a counter with the given name.
func New(name string) *counter.Counter

// NewStack returns a new stack counter with the given name and depth.
func NewStack(name string, depth int) *counter.StackCounter

// CountFlags creates a counter for every flag that is set
// and increments the counter. The name of the counter is
// the concatenation of prefix and the flag name.
func CountFlags(prefix string, flagSet flag.FlagSet)

// CountFlagValue creates a counter for the flag value
// if it is set and increments the counter. The name of the
// counter is the concatenation of prefix, the flagName, ":",
// and value.String() for the flag's value.
func CountFlagValue(prefix string, flagSet flag.FlagSet, flagName string)

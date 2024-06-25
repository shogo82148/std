// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cmd_go_bootstrap || compiler_bootstrap

package counter

import "github.com/shogo82148/std/flag"

func Open()
func Inc(name string)
func New(name string) dummyCounter
func NewStack(name string, depth int) dummyCounter
func CountFlags(name string, flagSet flag.FlagSet)
func CountFlagValue(prefix string, flagSet flag.FlagSet, flagName string)

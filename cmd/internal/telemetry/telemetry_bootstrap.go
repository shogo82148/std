// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cmd_go_bootstrap || compiler_bootstrap

package telemetry

import "github.com/shogo82148/std/flag"

func OpenCounters()
func MaybeParent()
func MaybeChild()
func Inc(name string)
func NewCounter(name string) dummyCounter
func NewStackCounter(name string, depth int) dummyCounter
func CountFlags(name string, flagSet flag.FlagSet)
func CountFlagValue(prefix string, flagSet flag.FlagSet, flagName string)
func Mode() string
func SetMode(mode string) error
func Dir() string

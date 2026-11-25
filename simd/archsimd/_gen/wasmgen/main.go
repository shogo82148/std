// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func CapitalizeFirst(s string) string

type OpFlags uint16

const (
	IsConst = OpFlags(1) << iota
	IsLoad
	IsStore
	IsShift
	IsSplat
	IsBitwise
	IsRelation
	IsTest
	IsExtract
	IsCommutative
	IsConversion
	NameHasFormat
	NonSigned
	EmulatedRule
	// Fixed permutation synthesized from i8x16.shuffle. Can take 1 or 2 args.
	// If 1 arg, the arg is duplicated to both inputs.
	IsSynthShuffle
)

func (o OpFlags) OneString() string

func (o OpFlags) String() string

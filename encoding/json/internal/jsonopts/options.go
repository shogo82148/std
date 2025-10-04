// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsonopts

import (
	"github.com/shogo82148/std/encoding/json/internal"
	"github.com/shogo82148/std/encoding/json/internal/jsonflags"
)

// Options is the common options type shared across json packages.
type Options interface {
	JSONOptions(internal.NotForPublicUse)
}

// Struct is the combination of all options in struct form.
// This is efficient to pass down the call stack and to query.
type Struct struct {
	Flags jsonflags.Flags

	CoderValues
	ArshalValues
}

type CoderValues struct {
	Indent       string
	IndentPrefix string
	ByteLimit    int64
	DepthLimit   int
}

type ArshalValues struct {
	Marshalers   any
	Unmarshalers any

	Format      string
	FormatDepth int
}

// DefaultOptionsV2 is the set of all options that define default v2 behavior.
var DefaultOptionsV2 = Struct{
	Flags: jsonflags.Flags{
		Presence: uint64(jsonflags.AllFlags & ^jsonflags.WhitespaceFlags),
		Values:   uint64(0),
	},
}

// DefaultOptionsV1 is the set of all options that define default v1 behavior.
var DefaultOptionsV1 = Struct{
	Flags: jsonflags.Flags{
		Presence: uint64(jsonflags.AllFlags & ^jsonflags.WhitespaceFlags),
		Values:   uint64(jsonflags.DefaultV1Flags),
	},
}

func (*Struct) JSONOptions(internal.NotForPublicUse)

// GetUnknownOption is injected by the "json" package to handle Options
// declared in that package so that "jsonopts" can handle them.
var GetUnknownOption = func(Struct, Options) (any, bool) { panic("unknown option") }

func GetOption[T any](opts Options, setter func(T) Options) (T, bool)

// JoinUnknownOption is injected by the "json" package to handle Options
// declared in that package so that "jsonopts" can handle them.
var JoinUnknownOption = func(Struct, Options) Struct { panic("unknown option") }

func (dst *Struct) Join(srcs ...Options)

func (dst *Struct) JoinWithoutCoderOptions(srcs ...Options)

type (
	Indent       string
	IndentPrefix string
	ByteLimit    int64
	DepthLimit   int
)

func (Indent) JSONOptions(internal.NotForPublicUse)
func (IndentPrefix) JSONOptions(internal.NotForPublicUse)
func (ByteLimit) JSONOptions(internal.NotForPublicUse)
func (DepthLimit) JSONOptions(internal.NotForPublicUse)

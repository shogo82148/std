// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ld

// A BuildMode indicates the sort of object we are building.
//
// Possible build modes are the same as those for the -buildmode flag
// in cmd/go, and are documented in 'go help buildmode'.
type BuildMode uint8

const (
	BuildModeUnset BuildMode = iota
	BuildModeExe
	BuildModePIE
	BuildModeCArchive
	BuildModeCShared
	BuildModeShared
	BuildModePlugin
)

// Set implements flag.Value to set the build mode based on the argument
// to the -buildmode flag.
func (mode *BuildMode) Set(s string) error

func (mode BuildMode) String() string

// LinkMode indicates whether an external linker is used for the final link.
type LinkMode uint8

const (
	LinkAuto LinkMode = iota
	LinkInternal
	LinkExternal
)

func (mode *LinkMode) Set(s string) error

func (mode *LinkMode) String() string

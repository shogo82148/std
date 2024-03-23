// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package buildcfg provides access to the build configuration
// described by the current environment. It is for use by build tools
// such as cmd/go or cmd/compile and for setting up go/build's Default context.
//
// Note that it does NOT provide access to the build configuration used to
// build the currently-running binary. For that, use runtime.GOOS etc
// as well as internal/goexperiment.
package buildcfg

import (
	"github.com/shogo82148/std/os"
)

var (
	GOROOT    = os.Getenv("GOROOT")
	GOARCH    = envOr("GOARCH", defaultGOARCH)
	GOOS      = envOr("GOOS", defaultGOOS)
	GO386     = envOr("GO386", defaultGO386)
	GOAMD64   = goamd64()
	GOARM     = goarm()
	GOARM64   = goarm64()
	GOMIPS    = gomips()
	GOMIPS64  = gomips64()
	GOPPC64   = goppc64()
	GORISCV64 = goriscv64()
	GOWASM    = gowasm()
	ToolTags  = toolTags()
	GO_LDSO   = defaultGO_LDSO
	Version   = version
)

// Error is one of the errors found (if any) in the build configuration.
var Error error

// Check exits the program with a fatal error if Error is non-nil.
func Check()

type Goarm64Features struct {
	Version string
	// Large Systems Extension
	LSE bool
	// ARM v8.0 Cryptographic Extension. It includes the following features:
	// * FEAT_AES, which includes the AESD and AESE instructions.
	// * FEAT_PMULL, which includes the PMULL, PMULL2 instructions.
	// * FEAT_SHA1, which includes the SHA1* instructions.
	// * FEAT_SHA256, which includes the SHA256* instructions.
	Crypto bool
}

func (g Goarm64Features) String() string

func ParseGoarm64(v string) (g Goarm64Features, e error)

// Returns true if g supports giving ARM64 ISA
// Note that this function doesn't accept / test suffixes (like ",lse" or ",crypto")
func (g Goarm64Features) Supports(s string) bool

func Getgoextlinkenabled() string

// GOGOARCH returns the name and value of the GO$GOARCH setting.
// For example, if GOARCH is "amd64" it might return "GOAMD64", "v2".
func GOGOARCH() (name, value string)

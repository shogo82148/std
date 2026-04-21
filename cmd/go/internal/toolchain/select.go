// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package toolchain implements dynamic switching of Go toolchains.
package toolchain

import (
	"github.com/shogo82148/std/cmd/go/internal/modload"
)

// FilterEnv returns a copy of env with internal GOTOOLCHAIN environment
// variables filtered out.
func FilterEnv(env []string) []string

// Select invokes a different Go toolchain if directed by
// the GOTOOLCHAIN environment variable or the user's configuration
// or go.mod file.
// It must be called early in startup.
// See https://go.dev/doc/toolchain#select.
func Select()

// TestVersionSwitch is set in the test go binary to the value in $TESTGO_VERSION_SWITCH.
// Valid settings are:
//
//	"switch" - simulate version switches by reinvoking the test go binary with a different TESTGO_VERSION.
//	"mismatch" - like "switch" but forget to set TESTGO_VERSION, so it looks like we invoked a mismatched toolchain
//	"loop" - like "mismatch" but forget the target check, causing a toolchain switching loop
var TestVersionSwitch string

// Exec invokes the specified Go toolchain or else prints an error and exits the process.
// If $GOTOOLCHAIN is set to path or min+path, Exec only considers the PATH
// as a source of Go toolchains. Otherwise Exec tries the PATH but then downloads
// a toolchain if necessary.
func Exec(s *modload.State, gotoolchain string)

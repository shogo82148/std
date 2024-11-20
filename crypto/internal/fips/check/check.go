// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package check implements the FIPS-140 load-time code+data verification.
// Every FIPS package providing cryptographic functionality except hmac and sha256
// must import crypto/internal/fips/check, so that the verification happens
// before initialization of package global variables.
// The hmac and sha256 packages are used by this package, so they cannot import it.
// Instead, those packages must be careful not to change global variables during init.
// (If necessary, we could have check call a PostCheck function in those packages
// after the check has completed.)
package check

import (
	"github.com/shogo82148/std/unsafe"
)

// Enabled reports whether verification was enabled.
// If Enabled returns true, then verification succeeded,
// because if it failed the binary would have panicked at init time.
func Enabled() bool

var Verified bool

// Supported reports whether the current GOOS/GOARCH is Supported at all.
func Supported() bool

// Linkinfo holds the go:fipsinfo symbol prepared by the linker.
// See cmd/link/internal/ld/fips.go for details.
//
//go:linkname Linkinfo go:fipsinfo
var Linkinfo struct {
	Magic [16]byte
	Sum   [32]byte
	Self  uintptr
	Sects [4]struct {
		Start unsafe.Pointer
		End   unsafe.Pointer
	}
}

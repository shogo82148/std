// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"github.com/shogo82148/std/io"
)

// Reader is an io.Reader that calls [drbg.Read].
//
// It should be used internally instead of [crypto/rand.Reader], because the
// latter can be set by applications outside of tests. These applications then
// risk breaking between Go releases, if the way the Reader is used changes.
var Reader io.Reader = reader{}

// SetTestingReader overrides all calls to [drbg.Read]. The Read method of
// r must never return an error or return short.
//
// SetTestingReader panics when building against Go Cryptographic Module v1.0.0.
//
// SetTestingReader is pulled by [testing/cryptotest.setGlobalRandom] via go:linkname.
//
//go:linkname SetTestingReader crypto/internal/rand.SetTestingReader
func SetTestingReader(r io.Reader)

// CustomReader returns [Reader] or, only if the GODEBUG setting
// "cryptocustomrand=1" is set, the provided io.Reader.
//
// If returning a non-default Reader, it calls [randutil.MaybeReadByte] on it.
func CustomReader(r io.Reader) io.Reader

// IsDefaultReader reports whether r is the default [crypto/rand.Reader].
//
// If true, the Read method of r can be assumed to call [drbg.Read].
func IsDefaultReader(r io.Reader) bool

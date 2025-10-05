// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package drbg provides cryptographically secure random bytes
// usable by FIPS code. In FIPS mode it uses an SP 800-90A Rev. 1
// Deterministic Random Bit Generator (DRBG). Otherwise,
// it uses the operating system's random number generator.
package drbg

import (
	"github.com/shogo82148/std/io"
)

// Read fills b with cryptographically secure random bytes. In FIPS mode, it
// uses an SP 800-90A Rev. 1 Deterministic Random Bit Generator (DRBG).
// Otherwise, it uses the operating system's random number generator.
func Read(b []byte)

// DefaultReader is a sentinel type, embedded in the default
// [crypto/rand.Reader], used to recognize it when passed to
// APIs that accept a rand io.Reader.
type DefaultReader interface{ defaultReader() }

// ReadWithReader uses Reader to fill b with cryptographically secure random
// bytes. It is intended for use in APIs that expose a rand io.Reader.
//
// If Reader is not the default Reader from crypto/rand,
// [randutil.MaybeReadByte] and [fips140.RecordNonApproved] are called.
func ReadWithReader(r io.Reader, b []byte) error

// ReadWithReaderDeterministic is like ReadWithReader, but it doesn't call
// [randutil.MaybeReadByte] on non-default Readers.
func ReadWithReaderDeterministic(r io.Reader, b []byte) error

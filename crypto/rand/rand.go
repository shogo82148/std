// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rand implements a cryptographically secure
// random number generator.
package rand

import (
	"github.com/shogo82148/std/crypto/internal/rand"
	"github.com/shogo82148/std/io"
)

// Reader is a global, shared instance of a cryptographically
// secure random number generator. It is safe for concurrent use.
//
//   - On Linux, FreeBSD, Dragonfly, and Solaris, Reader uses getrandom(2).
//   - On legacy Linux (< 3.17), Reader opens /dev/urandom on first use.
//   - On macOS, iOS, and OpenBSD Reader, uses arc4random_buf(3).
//   - On NetBSD, Reader uses the kern.arandom sysctl.
//   - On Windows, Reader uses the ProcessPrng API.
//   - On js/wasm, Reader uses the Web Crypto API.
//   - On wasip1/wasm, Reader uses random_get.
//
// In FIPS 140-3 mode, the output passes through an SP 800-90A Rev. 1
// Deterministric Random Bit Generator (DRBG).
var Reader io.Reader = rand.Reader

// Read fills b with cryptographically secure random bytes. It never returns an
// error, and always fills b entirely.
//
// Read calls [io.ReadFull] on [Reader] and crashes the program irrecoverably if
// an error is returned. The default Reader uses operating system APIs that are
// documented to never return an error on all but legacy Linux systems.
func Read(b []byte) (n int, err error)

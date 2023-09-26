// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rand implements a cryptographically secure
// random number generator.
package rand

import "github.com/shogo82148/std/io"

// Reader is a global, shared instance of a cryptographically
// secure random number generator.
//
// On Linux, FreeBSD, Dragonfly, NetBSD and Solaris, Reader uses getrandom(2) if
// available, /dev/urandom otherwise.
// On OpenBSD and macOS, Reader uses getentropy(2).
// On other Unix-like systems, Reader reads from /dev/urandom.
// On Windows systems, Reader uses the RtlGenRandom API.
// On JS/Wasm, Reader uses the Web Crypto API.
// On WASIP1/Wasm, Reader uses random_get from wasi_snapshot_preview1.
var Reader io.Reader

// Read is a helper function that calls Reader.Read using io.ReadFull.
// On return, n == len(b) if and only if err == nil.
func Read(b []byte) (n int, err error)

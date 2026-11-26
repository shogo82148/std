// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package scratch provides the scratch memory used by the FIPS 140-3 CPU
// jitter entropy source. It's outside the module boundary and the certified
// entropy source, so we can change how the memory is allocated.
package scratch

// Acquire returns a buffer of n bytes with unspecified contents, which must be
// passed to [Release] when no longer in use. n must be the same for all calls.
//
// Concurrent callers may get the same buffer, so the contents may change at
// any time, and must only be accessed with atomic operations.
func Acquire(n int) []byte

// Release releases a buffer returned by [Acquire].
// The buffer must not be used after it's released.
func Release(b []byte)

// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package drbg

// Read fills b with cryptographically secure random bytes. In FIPS mode, it
// uses an SP 800-90A Rev. 1 Deterministic Random Bit Generator (DRBG).
// Otherwise, it uses the operating system's random number generator.
func Read(b []byte)

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rand provides cryptographically secure random bytes from the
// operating system.
package sysrand

// Read fills b with cryptographically secure random bytes from the operating
// system. It always fills b entirely and crashes the program irrecoverably if
// an error is encountered. The operating system APIs are documented to never
// return an error on all but legacy Linux systems.
//
// Note that Read is not affected by [testing/cryptotest.SetGlobalRand], and it
// should not be used directly by algorithm implementations.
func Read(b []byte)

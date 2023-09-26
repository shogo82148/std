// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

// Unix cryptographically secure pseudorandom number
// generator.

package rand

// A reader satisfies reads by reading from urandomDevice

// altGetRandom if non-nil specifies an OS-specific function to get
// urandom-style randomness.

// hideAgainReader masks EAGAIN reads from /dev/urandom.
// See golang.org/issue/9205

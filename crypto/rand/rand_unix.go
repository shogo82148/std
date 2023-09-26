// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || nacl || netbsd || openbsd || plan9 || solaris
// +build darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris

// Unix cryptographically secure pseudorandom number
// generator.

package rand

// A devReader satisfies reads by reading the file named name.

// altGetRandom if non-nil specifies an OS-specific function to get
// urandom-style randomness.

// hideAgainReader masks EAGAIN reads from /dev/urandom.
// See golang.org/issue/9205

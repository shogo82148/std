// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || freebsd || linux || netbsd || openbsd || plan9
// +build darwin freebsd linux netbsd openbsd plan9

// Unix cryptographically secure pseudorandom number
// generator.

package rand

// A devReader satisfies reads by reading the file named name.

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || nacl || netbsd || openbsd || solaris
// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package runtime

// crashing is the number of m's we have waited for when implementing
// GOTRACEBACK=crash when a signal is received.

// testSigtrap is used by the runtime tests. If non-nil, it is called
// on SIGTRAP. If it returns true, the normal behavior on SIGTRAP is
// suppressed.

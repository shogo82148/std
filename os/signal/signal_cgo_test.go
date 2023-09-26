// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (darwin || dragonfly || freebsd || (linux && !android) || netbsd || openbsd) && cgo

// Note that this test does not work on Solaris: issue #22849.
// Don't run the test on Android because at least some versions of the
// C library do not define the posix_openpt function.

package signal_test

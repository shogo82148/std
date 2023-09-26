// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package runtime

// Stores the signal handlers registered before Go installed its own.
// These signal handlers will be invoked in cases where Go doesn't want to
// handle a particular signal (e.g., signal occurred on a non-Go thread).
// See sigfwdgo() for more information on when the signals are forwarded.
//
// Signal forwarding is currently available only on Darwin and Linux.

// sigmask represents a general signal mask compatible with the GOOS
// specific sigset types: the signal numbered x is represented by bit x-1
// to match the representation expected by sigprocmask.

// channels for synchronizing signal mask updates with the signal mask
// thread

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

package signal

// settleTime is an upper bound on how long we expect signals to take to be
// delivered. Lower values make the test faster, but also flakier — especially
// on heavily loaded systems.
//
// The current value is set based on flakes observed in the Go builders.

// fatalWaitingTime is an absurdly long time to wait for signals to be
// delivered but, using it, we (hopefully) eliminate test flakes on the
// build servers. See #46736 for discussion.

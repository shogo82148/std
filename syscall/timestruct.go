// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || (js && wasm) || linux || netbsd || openbsd || solaris

package syscall

// TimespecToNSec returns the time stored in ts as nanoseconds.
func TimespecToNsec(ts Timespec) int64

// NsecToTimespec converts a number of nanoseconds into a Timespec.
func NsecToTimespec(nsec int64) Timespec

// TimevalToNsec returns the time stored in tv as nanoseconds.
func TimevalToNsec(tv Timeval) int64

// NsecToTimeval converts a number of nanoseconds into a Timeval.
func NsecToTimeval(nsec int64) Timeval

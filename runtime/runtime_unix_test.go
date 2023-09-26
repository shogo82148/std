// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Only works on systems with syscall.Close.
// We need a fast system call to provoke the race,
// and Close(-1) is nearly universally fast.

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || plan9
// +build darwin dragonfly freebsd linux netbsd openbsd plan9

package runtime_test

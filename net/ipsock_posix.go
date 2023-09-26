// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || freebsd || linux || netbsd || openbsd || windows
// +build darwin freebsd linux netbsd openbsd windows

// Internet protocol family sockets for POSIX

package net

// A sockaddr represents a TCP, UDP or IP network address that can
// be converted into a syscall.Sockaddr.

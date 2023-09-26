// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js && !wasip1

package net

// Define a pair of blackholed (IPv4, IPv6) addresses, for which dialTCP is
// expected to hang until the timeout elapses. These addresses are reserved
// for benchmarking by RFC 6890.

// isEADDRINUSE reports whether err is syscall.EADDRINUSE.

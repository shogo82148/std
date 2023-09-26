// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || nacl || netbsd || openbsd || solaris
// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

// Read system port mappings from /etc/services

package net

// services contains minimal mappings between services names and port
// numbers for platforms that don't have a complete list of port numbers
// (some Solaris distros).

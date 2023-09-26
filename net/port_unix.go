// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || nacl
// +build darwin dragonfly freebsd linux netbsd openbsd solaris nacl

// Read system port mappings from /etc/services

package net

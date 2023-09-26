// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package net

// Test address from 192.0.2.0/24 block, reserved by RFC 5737 for documentation.
const TestAddr uint32 = 0xc0000201

// See RFC 6761 for further information about the reserved, pseudo
// domain names.

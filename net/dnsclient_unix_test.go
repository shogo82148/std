// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package net

// Test address from 192.0.2.0/24 block, reserved by RFC 5737 for documentation.
var TestAddr = [4]byte{0xc0, 0x00, 0x02, 0x01}

// Test address from 2001:db8::/32 block, reserved by RFC 3849 for documentation.
var TestAddr6 = [16]byte{0x20, 0x01, 0x0d, 0xb8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

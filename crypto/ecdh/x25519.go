// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecdh

// X25519 returns a Curve which implements the X25519 function over Curve25519
// (RFC 7748, Section 5).
//
// Multiple invocations of this function will return the same value, so it can
// be used for equality checks and switch statements.
func X25519() Curve

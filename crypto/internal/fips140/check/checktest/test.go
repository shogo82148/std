// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package checktest defines some code and data for use in
// the crypto/internal/fips140/check test.
package checktest

var NOPTRDATA int = 1

// The linkname here disables asan registration of this global,
// because asan gets mad about rodata globals.
//
//go:linkname RODATA crypto/internal/fips140/check/checktest.RODATA
var RODATA int32

// DATA needs to have both a pointer and an int so that _some_ of it gets
// initialized at link time, so it is treated as DATA and not BSS.
// The pointer is deferred to init time.
var DATA = struct {
	P *int
	X int
}{&NOPTRDATA, 3}

var NOPTRBSS int

var BSS *int

func TEXT()

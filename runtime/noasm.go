// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Routines that are implemented in assembly in asm_{amd64,386,arm,arm64,ppc64x,s390x}.s

//go:build mips64 || mips64le
// +build mips64 mips64le

package runtime

import _ "github.com/shogo82148/std/unsafe"

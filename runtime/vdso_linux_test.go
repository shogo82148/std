// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (386 || amd64)
// +build linux
// +build 386 amd64

package runtime_test

import (
	_ "unsafe"
)

//go:linkname __vdso_clock_gettime_sym runtime.__vdso_clock_gettime_sym

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (mips64 || mips64le)

package runtime

// see man 7 vdso : mips

// The symbol name is not __kernel_clock_gettime as suggested by the manpage;
// according to Linux source code it should be __vdso_clock_gettime instead.

// initialize to fall back to syscall

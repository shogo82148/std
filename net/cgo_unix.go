// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is called cgo_unix.go, but to allow syscalls-to-libc-based
// implementations to share the code, it does not use cgo directly.
// Instead of C.foo it uses _C_foo, which is defined in either
// cgo_unix_cgo.go or cgo_unix_syscall.go

//go:build !netgo && ((cgo && unix) || darwin)

package net

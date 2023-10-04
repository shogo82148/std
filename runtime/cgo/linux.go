// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Linux system call wrappers that provide POSIX semantics through the
// corresponding cgo->libc (nptl) wrappers for various system calls.

//go:build linux

package cgo

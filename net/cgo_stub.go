// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file holds stub versions of the cgo functions called on Unix systems.
// We build this file:
// - if using the netgo build tag on a Unix system
// - on a Unix system without the cgo resolver functions
//   (Darwin always provides the cgo functions, in cgo_unix_syscall.go)
// - on wasip1, where cgo is never available

//go:build (netgo && unix) || (unix && !cgo && !darwin) || js || wasip1

package net

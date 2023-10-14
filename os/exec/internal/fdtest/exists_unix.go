// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || wasm

// Package fdtest provides test helpers for working with file descriptors across exec.
package fdtest

// Exists returns true if fd is a valid file descriptor.
func Exists(fd uintptr) bool

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// mkasm_darwin.go generates assembly trampolines to call libSystem routines from Go.
// This program must be run after mksyscall.pl.
package main

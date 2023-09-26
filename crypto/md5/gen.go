// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// This program generates md5block.go
// Invoke as
//
//	go run gen.go [-full] -output md5block.go
//
// The -full flag causes the generated code to do a full
// (16x) unrolling instead of a 4x unrolling.

package main

type Data struct {
	a, b, c, d string
	Shift1     []int
	Shift2     []int
	Shift3     []int
	Shift4     []int
	Table1     []uint32
	Table2     []uint32
	Table3     []uint32
	Table4     []uint32
	Full       bool
}

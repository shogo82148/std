// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// This program generates md5block.go
// Invoke as
//
//	go run gen.go -output md5block.go

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
}

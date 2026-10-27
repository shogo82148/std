// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test cases for cgo.
// Both the import "C" prologue and the main file are sorted by issue number.
// This file contains //export directives on Go functions
// and so it must NOT contain C definitions (only declarations).
// See test.go for C definitions.

package cgotest

import (
	"github.com/shogo82148/std/unsafe"
)

import "C"

//export ReturnIntLong
func ReturnIntLong() (int, C.long)

//export Add
func Add(x int)

//export BackIntoGo
func BackIntoGo()

//export Issue1560FromC
func Issue1560FromC()

func Issue1560FromGo()

//export GoIssue6833Func
func GoIssue6833Func(aui uint, aui64 uint64) uint64

const CString = "C string"

//export CheckIssue6907Go
func CheckIssue6907Go(s string) C.int

//export Test8945
func Test8945()

//export GoFunc37033
func GoFunc37033(handle C.uintptr_t)

// issue 38408
// A typedef pointer can be used as the element type.
// No runtime test; just make sure it compiles.
var _ C.PIssue38408 = &C.Issue38408{i: 1}

//export Issue44648A
func Issue44648A(issue44648F C.issue44648F, data unsafe.Pointer) C.int

//export Issue44648B
func Issue44648B(issue44648T C.issue44648T) (C.issue44648T, C.issue44648T)

//export Issue44648C
func Issue44648C(r C.int) (a, b C.int)

//export GoFunc49633
func GoFunc49633(context unsafe.Pointer)

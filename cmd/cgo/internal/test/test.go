// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test cases for cgo.
// Both the import "C" prologue and the main file are sorted by issue number.
// This file contains C definitions (not just declarations)
// and so it must NOT contain any //export directives on Go functions.
// See testx.go for exports.

package cgotest

import "C"

import (
	"github.com/shogo82148/std/testing"
	"github.com/shogo82148/std/unsafe"
)

const EINVAL = C.EINVAL

var KILO = C.KILO

func Strtol(s string, base int) (int, error)

func Atol(s string) int

type Context struct {
	ctx *C.struct_ibv_context
}

func TestComplexAlign(t *testing.T)

// issue 1222
type AsyncEvent struct {
	event C.struct_ibv_async_event
}

var _ = C.struct_issue8428one{
	b: C.char(0),
}

var _ = C.struct_issue8428two{
	p:    unsafe.Pointer(nil),
	b:    C.char(0),
	rest: [0]C.char{},
}

var _ = C.struct_issue8428three{
	w: [1][2][3][0]C.char{},
	x: [2][3][0][1]C.char{},
	y: [3][0][1][2]C.char{},
	z: [0][1][2][3]C.char{},
}

var _ C.complexfloat
var _ C.complexdouble

var _, _ = C.abs(0)

// issue 22958
// Nothing to run, just make sure this compiles.
var Vissue22958 C.issue22958Type

func Issue23720F()

func Issue29383(n, size uint) int

var Vissue29748 = C.f29748(&C.S29748{
	nil,
})

func Fissue299748()

func Issue31093()

func Issue40494()

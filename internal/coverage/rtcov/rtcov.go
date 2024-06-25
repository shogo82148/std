// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtcov

import "github.com/shogo82148/std/unsafe"

// CovMetaBlob is a container for holding the meta-data symbol (an
// RODATA variable) for an instrumented Go package. Here "p" points to
// the symbol itself, "len" is the length of the sym in bytes, and
// "hash" is an md5sum for the sym computed by the compiler. When
// the init function for a coverage-instrumented package executes, it
// will make a call into the runtime which will create a covMetaBlob
// object for the package and chain it onto a global list.
type CovMetaBlob struct {
	P                  *byte
	Len                uint32
	Hash               [16]byte
	PkgPath            string
	PkgID              int
	CounterMode        uint8
	CounterGranularity uint8
}

// CovCounterBlob is a container for encapsulating a counter section
// (BSS variable) for an instrumented Go module. Here "counters"
// points to the counter payload and "len" is the number of uint32
// entries in the section.
type CovCounterBlob struct {
	Counters *uint32
	Len      uint64
}

// Meta is the top-level container for bits of state related to
// code coverage meta-data in the runtime.
var Meta struct {
	List []CovMetaBlob

	PkgMap map[int]int

	hardCodedListNeedsUpdating bool
}

// AddMeta is invoked during package "init" functions by the
// compiler when compiling for coverage instrumentation; here 'p' is a
// meta-data blob of length 'dlen' for the package in question, 'hash'
// is a compiler-computed md5.sum for the blob, 'pkpath' is the
// package path, 'pkid' is the hard-coded ID that the compiler is
// using for the package (or -1 if the compiler doesn't think a
// hard-coded ID is needed), and 'cmode'/'cgran' are the coverage
// counter mode and granularity requested by the user. Return value is
// the ID for the package for use by the package code itself,
// or 0 for impossible errors.
func AddMeta(p unsafe.Pointer, dlen uint32, hash [16]byte, pkgpath string, pkgid int, cmode uint8, cgran uint8) uint32

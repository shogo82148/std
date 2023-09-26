// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

// Generate the constant table associated with the poly used by the
// vpmsumd crc32 algorithm.
//
// go run gen_const_ppc64le.go
//
// generates crc32_table_ppc64le.s

// The following is derived from code written by Anton Blanchard
// <anton@au.ibm.com> found at https://github.com/antonblanchard/crc32-vpmsum.
// The original is dual licensed under GPL and Apache 2.  As the copyright holder
// for the work, IBM has contributed this new work under the golang license.

// This code was written in Go based on the original C implementation.

// This is a tool needed to generate the appropriate constants needed for
// the vpmsum algorithm.  It is included to generate new constant tables if
// new polynomial values are included in the future.

package main

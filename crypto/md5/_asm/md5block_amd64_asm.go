// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Original source:
//	http://www.zorinaq.com/papers/md5-amd64.html
//	http://www.zorinaq.com/papers/md5-amd64.tar.bz2
//
// Translated from Perl generating GNU assembly into
// #defines generating 6a assembly by the Go Authors.

package main

func ROUND1(a, b, c, d GPPhysical, index int, konst, shift uint64)

// Uses https://github.com/animetosho/md5-optimisation#dependency-shortcut-in-g-function
func ROUND2(a, b, c, d GPPhysical, index int, konst, shift uint64)

// Uses https://github.com/animetosho/md5-optimisation#h-function-re-use
func ROUND3FIRST(a, b, c, d GPPhysical, index int, konst, shift uint64)

func ROUND3(a, b, c, d GPPhysical, index int, konst, shift uint64)

func ROUND4(a, b, c, d GPPhysical, index int, konst, shift uint64)

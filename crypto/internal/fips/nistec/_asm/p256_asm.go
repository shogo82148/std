// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains constant-time, 64-bit assembly implementation of
// P256. The optimizations performed here are described in detail in:
// S.Gueron and V.Krasnov, "Fast prime field elliptic-curve cryptography with
//                          256-bit primes"
// https://link.springer.com/article/10.1007%2Fs13389-014-0090-x
// https://eprint.iacr.org/2013/816.pdf

package main

type MemFunc func(off int) Mem

func LDacc(src MemFunc)

func LDt(src MemFunc)

func ST(dst MemFunc)

func STt(dst MemFunc)

const ThatPeskyUnicodeDot = "\u00b7"

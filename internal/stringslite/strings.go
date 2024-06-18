// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package stringslite implements a subset of strings,
// only using packages that may be imported by "os".
//
// Tests for these functions are in the strings package.
package stringslite

func HasPrefix(s, prefix string) bool

func HasSuffix(s, suffix string) bool

func IndexByte(s string, c byte) int

func Index(s, substr string) int

func Cut(s, sep string) (before, after string, found bool)

func CutPrefix(s, prefix string) (after string, found bool)

func CutSuffix(s, suffix string) (before string, found bool)

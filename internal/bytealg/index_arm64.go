// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytealg

// bytes.Index と strings.Index は、この長さ以下の haystack を、
// 先に IndexByte で先読みせずに、そのまま Index に渡します。
const MaxBruteForce = 64

// Cutover reports the number of failures of IndexByte we should tolerate
// before switching over to Index.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
func Cutover(n int) int

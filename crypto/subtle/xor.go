// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package subtle

<<<<<<< HEAD
// XORBytesはdst[i] = x[i] ^ y[i]をi < n = min(len(x), len(y))のすべてのiに対して設定し、
// dstに書き込まれたバイト数であるnを返します。
// もしdstの長さが少なくともnでない場合、
// XORBytesはなにもdstに書き込まずにパニックを起こします。
=======
// XORBytes sets dst[i] = x[i] ^ y[i] for all i < n = min(len(x), len(y)),
// returning n, the number of bytes written to dst.
//
// If dst does not have length at least n,
// XORBytes panics without writing anything to dst.
//
// dst and x or y may overlap exactly or not at all,
// otherwise XORBytes may panic.
>>>>>>> upstream/release-branch.go1.25
func XORBytes(dst, x, y []byte) int

// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package subtle

// XORBytesはdst[i] = x[i] ^ y[i]をi < n = min(len(x), len(y))のすべてのiに対して設定し、
// dstに書き込まれたバイト数であるnを返します。
// もしdstの長さが少なくともnでない場合、
// XORBytesはなにもdstに書き込まずにパニックを起こします。
func XORBytes(dst, x, y []byte) int

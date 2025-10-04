// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package subtle

// XORBytesはすべてのi < n = min(len(x), len(y))に対してdst[i] = x[i] ^ y[i]を設定し、
// dstに書き込まれたバイト数nを返します。
//
// dstの長さが少なくともnでない場合、
// XORBytesはdstに何も書き込まずにパニックします。
//
// dstとxまたはyは完全に重複するか全く重複しないかのいずれかである必要があり、
// そうでない場合XORBytesはパニックすることがあります。
func XORBytes(dst, x, y []byte) int

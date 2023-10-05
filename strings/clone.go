// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

// Cloneは、sの新しいコピーを返します。
// sを新しい割り当てにコピーすることを保証します。
// これは、大きな文字列の小さなサブストリングのみを保持する場合に重要な場合があります。
// Cloneを使用することで、このようなプログラムがより少ないメモリを使用できるようになります。
// もちろん、Cloneを使用するとコピーが作成されるため、Cloneの過剰使用はプログラムのメモリ使用量を増やす可能性があります。
// Cloneは通常、プロファイリングによって必要であることが示された場合にのみ使用する必要があります。
// 長さがゼロの文字列の場合、文字列 "" が返され、割り当ては行われません。
func Clone(s string) string

// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scan

// FilterNilRVV は、bufp[0:n] の nil でない値を bufp の先頭に詰め、
// 詰めた後の長さを返します。V 拡張が利用できる場合は RVV 実装を使用し、
// そうでない場合は汎用の FilterNil にフォールバックします。
func FilterNilRVV(bufp *uintptr, n int32) int32

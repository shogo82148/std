// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecdh

// X25519はCurve25519上でX25519関数を実装するCurveを返します
// (RFC 7748、セクション5)。
//
// この関数の複数の呼び出しは同じ値を返すため、等価性チェックやスイッチ文に使用できます。
func X25519() Curve

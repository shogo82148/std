// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecdh

// P256は、NIST P-256 (FIPS 186-3, セクション D.2.3)、またはsecp256r1またはprime256v1としても知られる曲線を実装する [Curve] を返します。
//
// この関数の複数の呼び出しは、等値チェックやスイッチ文に使用できる同じ値を返します。
func P256() Curve

// P384は、NIST P-384（FIPS 186-3、セクション D.2.4）またはsecp384r1としても知られる
// 曲線を実装する [Curve] を返します。
//
// この関数の複数回呼び出しでは、同じ値が返され、等価性のチェックやswitch文に使用できます。
func P384() Curve

// P521は、NIST P-521（FIPS 186-3、セクションD.2.5）で定義されている、secp521r1としても知られる曲線を実装する [Curve] を返します。
//
// この関数の複数の呼び出しは、同じ値を返します。これは、等値比較やスイッチ文で使用することができます。
func P521() Curve

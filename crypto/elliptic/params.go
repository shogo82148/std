// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elliptic

import "github.com/shogo82148/std/math/big"

// CurveParamsは楕円曲線のパラメータを含み、Curveの汎用で非定数時間の実装も提供します。
//
// 汎用のCurve実装は非推奨であり、カスタム曲線（P224（）、P256（）、P384（）、P521（）によって返されない曲線）を使用することは、
// いかなるセキュリティプロパティも保証しないことに注意してください。
type CurveParams struct {
	P       *big.Int
	N       *big.Int
	B       *big.Int
	Gx, Gy  *big.Int
	BitSize int
	Name    string
}

func (curve *CurveParams) Params() *CurveParams

// IsOnCurveはCurve.IsOnCurveを実装します。
//
// Deprecated: CurveParamsのメソッドは非推奨であり、任意のセキュリティプロパティを提供することは保証されていません。
// ECDHにはcrypto/ecdhパッケージを使用してください。
// ECDSAには、P224()、P256()、P384()、またはP521()から直接返されるCurve値を使用して、crypto/ecdsaパッケージを使用してください。
func (curve *CurveParams) IsOnCurve(x, y *big.Int) bool

// AddはCurve.Addを実装します。
//
// 廃止予定: CurveParamsのメソッドは廃止予定であり、何のセキュリティプロパティも保証されません。
// ECDHについては、crypto/ecdhパッケージを使用してください。
// ECDSAについては、直接P224()、P256()、P384()、またはP521()から返されるCurve値と一緒にcrypto/ecdsaパッケージを使用してください。
func (curve *CurveParams) Add(x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int)

// DoubleはCurve.Doubleを実装します。
//
// 廃止予定: CurveParamsメソッドは廃止予定であり、セキュリティプロパティを提供することを保証しません。
// ECDHにはcrypto/ecdhパッケージを使用してください。
// ECDSAには、直接P224()、P256()、P384()、またはP521()から返されるCurveの値を使用してcrypto/ecdsaパッケージを使用してください。
func (curve *CurveParams) Double(x1, y1 *big.Int) (*big.Int, *big.Int)

// ScalarMultはCurve.ScalarMultを実装します。
//
// Deprecated: CurveParamsのメソッドは非推奨であり、いかなるセキュリティ保護も提供しません。
// ECDHの場合は、crypto/ecdhパッケージを使用してください。
// ECDSAの場合は、crypto/ecdsaパッケージを使用し、直接Curve値をP224()、P256()、P384()、またはP521()から取得してください。
func (curve *CurveParams) ScalarMult(Bx, By *big.Int, k []byte) (*big.Int, *big.Int)

// ScalarBaseMultはCurve.ScalarBaseMultを実装します。
//
// 廃止予定: CurveParamsメソッドは廃止されており、安全性を保証するものではありません。
// ECDHにはcrypto/ecdhパッケージを使用してください。
// ECDSAにはP224（）、P256（）、P384（）またはP521（）から直接返されるCurve値を使用して、crypto/ecdsaパッケージを使用してください。
func (curve *CurveParams) ScalarBaseMult(k []byte) (*big.Int, *big.Int)

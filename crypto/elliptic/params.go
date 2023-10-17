// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elliptic

import "github.com/shogo82148/std/math/big"

// CurveParamsは楕円曲線のパラメータを含み、 [Curve] の汎用で非定数時間の実装も提供します。
//
// 汎用のCurve実装は非推奨であり、カスタム曲線（ [P224] 、 [P256] 、 [P384] 、 [P521] によって返されない曲線）を使用することは、
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

// IsOnCurveは [Curve.IsOnCurve] を実装します。
//
// Deprecated: [CurveParams] のメソッドは非推奨であり、任意のセキュリティプロパティを提供することは保証されていません。
// ECDHには [crypto/ecdh] パッケージを使用してください。
// ECDSAには、 [P224] 、 [P256] 、 [P384] 、または [P521] から直接返されるCurve値を使用して、 [crypto/ecdsa] パッケージを使用してください。
func (curve *CurveParams) IsOnCurve(x, y *big.Int) bool

// Addは [Curve.Add] を実装します。
//
// Deprecated: [CurveParams] のメソッドは非推奨であり、任意のセキュリティプロパティを提供することは保証されていません。
// ECDHには [crypto/ecdh] パッケージを使用してください。
// ECDSAには、 [P224] 、 [P256] 、 [P384] 、または [P521] から直接返されるCurve値を使用して、 [crypto/ecdsa] パッケージを使用してください。
func (curve *CurveParams) Add(x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int)

// Doubleは [Curve.Double] を実装します。
//
// Deprecated: [CurveParams] のメソッドは非推奨であり、任意のセキュリティプロパティを提供することは保証されていません。
// ECDHには [crypto/ecdh] パッケージを使用してください。
// ECDSAには、 [P224] 、 [P256] 、 [P384] 、または [P521] から直接返されるCurve値を使用して、 [crypto/ecdsa] パッケージを使用してください。
func (curve *CurveParams) Double(x1, y1 *big.Int) (*big.Int, *big.Int)

// ScalarMultは [Curve.ScalarMult] を実装します。
//
// Deprecated: [CurveParams] のメソッドは非推奨であり、任意のセキュリティプロパティを提供することは保証されていません。
// ECDHには [crypto/ecdh] パッケージを使用してください。
// ECDSAには、 [P224] 、 [P256] 、 [P384] 、または [P521] から直接返されるCurve値を使用して、 [crypto/ecdsa] パッケージを使用してください。
func (curve *CurveParams) ScalarMult(Bx, By *big.Int, k []byte) (*big.Int, *big.Int)

// ScalarBaseMultは [Curve.ScalarBaseMult] を実装します。
//
// Deprecated: [CurveParams] のメソッドは非推奨であり、任意のセキュリティプロパティを提供することは保証されていません。
// ECDHには [crypto/ecdh] パッケージを使用してください。
// ECDSAには、 [P224] 、 [P256] 、 [P384] 、または [P521] から直接返されるCurve値を使用して、 [crypto/ecdsa] パッケージを使用してください。
func (curve *CurveParams) ScalarBaseMult(k []byte) (*big.Int, *big.Int)

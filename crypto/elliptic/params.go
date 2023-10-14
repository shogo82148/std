// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elliptic

import "github.com/shogo82148/std/math/big"

// CurveParams contains the parameters of an elliptic curve and also provides
// a generic, non-constant time implementation of [Curve].
//
// The generic Curve implementation is deprecated, and using custom curves
// (those not returned by [P224], [P256], [P384], and [P521]) is not guaranteed
// to provide any security property.
type CurveParams struct {
	P       *big.Int
	N       *big.Int
	B       *big.Int
	Gx, Gy  *big.Int
	BitSize int
	Name    string
}

func (curve *CurveParams) Params() *CurveParams

// IsOnCurve implements [Curve.IsOnCurve].
//
// Deprecated: the [CurveParams] methods are deprecated and are not guaranteed to
// provide any security property. For ECDH, use the [crypto/ecdh] package.
// For ECDSA, use the [crypto/ecdsa] package with a [Curve] value returned directly
// from [P224], [P256], [P384], or [P521].
func (curve *CurveParams) IsOnCurve(x, y *big.Int) bool

// Add implements [Curve.Add].
//
// Deprecated: the [CurveParams] methods are deprecated and are not guaranteed to
// provide any security property. For ECDH, use the [crypto/ecdh] package.
// For ECDSA, use the [crypto/ecdsa] package with a [Curve] value returned directly
// from [P224], [P256], [P384], or [P521].
func (curve *CurveParams) Add(x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int)

// Double implements [Curve.Double].
//
// Deprecated: the [CurveParams】 methods are deprecated and are not guaranteed to
// provide any security property. For ECDH, use the [crypto/ecdh] package.
// For ECDSA, use the [crypto/ecdsa] package with a [Curve] value returned directly
// from [P224], [P256], [P384], or [P521].
func (curve *CurveParams) Double(x1, y1 *big.Int) (*big.Int, *big.Int)

// ScalarMult implements [Curve.ScalarMult].
//
// Deprecated: the [CurveParams] methods are deprecated and are not guaranteed to
// provide any security property. For ECDH, use the [crypto/ecdh] package.
// For ECDSA, use the [crypto/ecdsa] package with a [Curve] value returned directly
// from [P224], [P256], [P384], or [P521].
func (curve *CurveParams) ScalarMult(Bx, By *big.Int, k []byte) (*big.Int, *big.Int)

// ScalarBaseMult implements [Curve.ScalarBaseMult].
//
// Deprecated: the [CurveParams] methods are deprecated and are not guaranteed to
// provide any security property. For ECDH, use the [crypto/ecdh] package.
// For ECDSA, use the [crypto/ecdsa] package with a [Curve] value returned directly
// from [P224], [P256], [P384], or [P521].
func (curve *CurveParams) ScalarBaseMult(k []byte) (*big.Int, *big.Int)

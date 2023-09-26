// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elliptic

import "github.com/shogo82148/std/math/big"

// CurveParams contains the parameters of an elliptic curve and also provides
// a generic, non-constant time implementation of Curve.
type CurveParams struct {
	P       *big.Int
	N       *big.Int
	B       *big.Int
	Gx, Gy  *big.Int
	BitSize int
	Name    string
}

func (curve *CurveParams) Params() *CurveParams

func (curve *CurveParams) IsOnCurve(x, y *big.Int) bool

func (curve *CurveParams) Add(x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int)

func (curve *CurveParams) Double(x1, y1 *big.Int) (*big.Int, *big.Int)

func (curve *CurveParams) ScalarMult(Bx, By *big.Int, k []byte) (*big.Int, *big.Int)

func (curve *CurveParams) ScalarBaseMult(k []byte) (*big.Int, *big.Int)

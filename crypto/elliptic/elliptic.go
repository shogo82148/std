// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package elliptic implements several standard elliptic curves over prime
// fields.
package elliptic

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// A Curve represents a short-form Weierstrass curve with a=-3.
// See https://www.hyperelliptic.org/EFD/g1p/auto-shortw.html
type Curve interface {
	Params() *CurveParams

	IsOnCurve(x, y *big.Int) bool

	Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int)

	Double(x1, y1 *big.Int) (x, y *big.Int)

	ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int)

	ScalarBaseMult(k []byte) (x, y *big.Int)
}

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

// GenerateKey returns a public/private key pair. The private key is
// generated using the given reader, which must return random data.
func GenerateKey(curve Curve, rand io.Reader) (priv []byte, x, y *big.Int, err error)

// Marshal converts a point into the uncompressed form specified in section 4.3.6 of ANSI X9.62.
func Marshal(curve Curve, x, y *big.Int) []byte

// Unmarshal converts a point, serialized by Marshal, into an x, y pair.
// It is an error if the point is not in uncompressed form or is not on the curve.
// On error, x = nil.
func Unmarshal(curve Curve, data []byte) (x, y *big.Int)

// P256 returns a Curve which implements NIST P-256 (FIPS 186-3, section D.2.3),
// also known as secp256r1 or prime256v1. The CurveParams.Name of this Curve is
// "P-256".
//
// Multiple invocations of this function will return the same value, so it can
// be used for equality checks and switch statements.
//
// The cryptographic operations are implemented using constant-time algorithms.
func P256() Curve

// P384 returns a Curve which implements NIST P-384 (FIPS 186-3, section D.2.4),
// also known as secp384r1. The CurveParams.Name of this Curve is "P-384".
//
// Multiple invocations of this function will return the same value, so it can
// be used for equality checks and switch statements.
//
// The cryptographic operations do not use constant-time algorithms.
func P384() Curve

// P521 returns a Curve which implements NIST P-521 (FIPS 186-3, section D.2.5),
// also known as secp521r1. The CurveParams.Name of this Curve is "P-521".
//
// Multiple invocations of this function will return the same value, so it can
// be used for equality checks and switch statements.
//
// The cryptographic operations do not use constant-time algorithms.
func P521() Curve

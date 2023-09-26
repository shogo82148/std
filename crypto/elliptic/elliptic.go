// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package elliptic implements the standard NIST P-224, P-256, P-384, and P-521
// elliptic curves over prime fields.
//
// The P224(), P256(), P384() and P521() values are necessary to use the crypto/ecdsa package.
// Most other uses should migrate to the more efficient and safer crypto/ecdh package.
package elliptic

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// A Curve represents a short-form Weierstrass curve with a=-3.
//
// The behavior of Add, Double, and ScalarMult when the input is not a point on
// the curve is undefined.
//
// Note that the conventional point at infinity (0, 0) is not considered on the
// curve, although it can be returned by Add, Double, ScalarMult, or
// ScalarBaseMult (but not the Unmarshal or UnmarshalCompressed functions).
type Curve interface {
	Params() *CurveParams

	IsOnCurve(x, y *big.Int) bool

	Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int)

	Double(x1, y1 *big.Int) (x, y *big.Int)

	ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int)

	ScalarBaseMult(k []byte) (x, y *big.Int)
}

// GenerateKey returns a public/private key pair. The private key is
// generated using the given reader, which must return random data.
//
// Note: for ECDH, use the GenerateKey methods of the crypto/ecdh package;
// for ECDSA, use the GenerateKey function of the crypto/ecdsa package.
func GenerateKey(curve Curve, rand io.Reader) (priv []byte, x, y *big.Int, err error)

// Marshal converts a point on the curve into the uncompressed form specified in
// SEC 1, Version 2.0, Section 2.3.3. If the point is not on the curve (or is
// the conventional point at infinity), the behavior is undefined.
//
// Note: for ECDH, use the crypto/ecdh package. This function returns an
// encoding equivalent to that of PublicKey.Bytes in crypto/ecdh.
func Marshal(curve Curve, x, y *big.Int) []byte

// MarshalCompressed converts a point on the curve into the compressed form
// specified in SEC 1, Version 2.0, Section 2.3.3. If the point is not on the
// curve (or is the conventional point at infinity), the behavior is undefined.
func MarshalCompressed(curve Curve, x, y *big.Int) []byte

// unmarshaler is implemented by curves with their own constant-time Unmarshal.
//
// There isn't an equivalent interface for Marshal/MarshalCompressed because
// that doesn't involve any mathematical operations, only FillBytes and Bit.

// Assert that the known curves implement unmarshaler.
var _ = []unmarshaler{p224, p256, p384, p521}

// Unmarshal converts a point, serialized by Marshal, into an x, y pair. It is
// an error if the point is not in uncompressed form, is not on the curve, or is
// the point at infinity. On error, x = nil.
//
// Note: for ECDH, use the crypto/ecdh package. This function accepts an
// encoding equivalent to that of the NewPublicKey methods in crypto/ecdh.
func Unmarshal(curve Curve, data []byte) (x, y *big.Int)

// UnmarshalCompressed converts a point, serialized by MarshalCompressed, into
// an x, y pair. It is an error if the point is not in compressed form, is not
// on the curve, or is the point at infinity. On error, x = nil.
func UnmarshalCompressed(curve Curve, data []byte) (x, y *big.Int)

// P224 returns a Curve which implements NIST P-224 (FIPS 186-3, section D.2.2),
// also known as secp224r1. The CurveParams.Name of this Curve is "P-224".
//
// Multiple invocations of this function will return the same value, so it can
// be used for equality checks and switch statements.
//
// The cryptographic operations are implemented using constant-time algorithms.
func P224() Curve

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
// The cryptographic operations are implemented using constant-time algorithms.
func P384() Curve

// P521 returns a Curve which implements NIST P-521 (FIPS 186-3, section D.2.5),
// also known as secp521r1. The CurveParams.Name of this Curve is "P-521".
//
// Multiple invocations of this function will return the same value, so it can
// be used for equality checks and switch statements.
//
// The cryptographic operations are implemented using constant-time algorithms.
func P521() Curve

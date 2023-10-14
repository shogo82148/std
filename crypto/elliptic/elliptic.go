// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージellipticは、素数体上の標準的なNIST P-224、P-256、P-384、およびP-521楕円曲線を実装しています。
//
// このパッケージの直接使用は非推奨であり、[P224]、[P256]、[P384]、[P521]の値は[crypto/ecdsa]を使用するために必要です。その他のほとんどの用途は、効率的かつ安全な[crypto/ecdh]または低レベルな機能のためのサードパーティのモジュールに移行する必要があります。
package elliptic

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// Curveはa=-3の短形式Weierstrass曲線を表します。
//
// 入力が曲線上の点でない場合、Add、Double、およびScalarMultの動作は未定義です。
//
<<<<<<< HEAD
// 0, 0のような通常の無限遠点は、曲線上には考慮されていませんが、Add、Double、
// ScalarMult、またはScalarBaseMultで返される場合があります
// （ただし、UnmarshalまたはUnmarshalCompressed関数では返されません）。
//
// P224（）、P256（）、P384（）、およびP521（）以外のCurve実装を使用することは非推奨です。
=======
// Note that the conventional point at infinity (0, 0) is not considered on the
// curve, although it can be returned by Add, Double, ScalarMult, or
// ScalarBaseMult (but not the [Unmarshal] or [UnmarshalCompressed] functions).
//
// Using Curve implementations besides those returned by [P224], [P256], [P384],
// and [P521] is deprecated.
>>>>>>> upstream/master
type Curve interface {
	// Params returns the parameters for the curve.
	Params() *CurveParams

	// IsOnCurve reports whether the given (x,y) lies on the curve.
	//
	// Deprecated: this is a low-level unsafe API. For ECDH, use the crypto/ecdh
	// package. The NewPublicKey methods of NIST curves in crypto/ecdh accept
	// the same encoding as the Unmarshal function, and perform on-curve checks.
	IsOnCurve(x, y *big.Int) bool

	// Add returns the sum of (x1,y1) and (x2,y2).
	//
	// Deprecated: this is a low-level unsafe API.
	Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int)

	// Double returns 2*(x,y).
	//
	// Deprecated: this is a low-level unsafe API.
	Double(x1, y1 *big.Int) (x, y *big.Int)

	// ScalarMult returns k*(x,y) where k is an integer in big-endian form.
	//
	// Deprecated: this is a low-level unsafe API. For ECDH, use the crypto/ecdh
	// package. Most uses of ScalarMult can be replaced by a call to the ECDH
	// methods of NIST curves in crypto/ecdh.
	ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int)

	// ScalarBaseMult returns k*G, where G is the base point of the group
	// and k is an integer in big-endian form.
	//
	// Deprecated: this is a low-level unsafe API. For ECDH, use the crypto/ecdh
	// package. Most uses of ScalarBaseMult can be replaced by a call to the
	// PrivateKey.PublicKey method in crypto/ecdh.
	ScalarBaseMult(k []byte) (x, y *big.Int)
}

<<<<<<< HEAD
// GenerateKeyは公開鍵と秘密鍵のペアを生成します。秘密鍵は与えられたリーダーを使用して生成されますが、ランダムデータを返す必要があります。
// 廃止予定: ECDHの場合はcrypto/ecdhパッケージのGenerateKeyメソッドを使用してください。
// ECDSAの場合はcrypto/ecdsaパッケージのGenerateKey関数を使用してください。
=======
// GenerateKey returns a public/private key pair. The private key is
// generated using the given reader, which must return random data.
//
// Deprecated: for ECDH, use the GenerateKey methods of the [crypto/ecdh] package;
// for ECDSA, use the GenerateKey function of the crypto/ecdsa package.
>>>>>>> upstream/master
func GenerateKey(curve Curve, rand io.Reader) (priv []byte, x, y *big.Int, err error)

// Marshalは、曲線上の点を、SEC 1、バージョン2.0、セクション2.3.3で指定された非圧縮形式に変換します。もし点が曲線上にない場合（または通常の無限遠点の場合）、動作は未定義です。
//
// 廃止されました：ECDHには、crypto/ecdhパッケージを使用してください。この関数は、crypto/ecdhのPublicKey.Bytesと同等のエンコーディングを返します。
func Marshal(curve Curve, x, y *big.Int) []byte

// MarshalCompressedは、曲線上の点をSEC 1、バージョン2.0、セクション2.3.3で指定された圧縮形式に変換します。点が曲線上にない場合（または無限遠点である場合）、動作は未定義です。
func MarshalCompressed(curve Curve, x, y *big.Int) []byte

// 既知の曲線がunmarshalerを実装していることを確認します。
var _ = []unmarshaler{p224, p256, p384, p521}

<<<<<<< HEAD
// UnmarshalはMarshalによってシリアライズされたポイントをx、yのペアに変換します。非圧縮形式でない場合、曲線上にない場合、または無限遠点の場合はエラーです。エラーの場合、x = nilです。
// 廃止予定：ECDHでは、crypto/ecdhパッケージを使用してください。この関数は、crypto/ecdhのNewPublicKeyメソッドで使用されるエンコーディングと同等のエンコーディングを受け入れます。
func Unmarshal(curve Curve, data []byte) (x, y *big.Int)

// UnmarshalCompressedはMarshalCompressedによって直列化された点を、xとyの組へと変換します。
// 圧縮形式でない場合、曲線上にない場合、または無限遠点の場合はエラーです。 エラー時には、x = nil です。
func UnmarshalCompressed(curve Curve, data []byte) (x, y *big.Int)

// P224は、NIST P-224（FIPS 186-3、セクションD.2.2）で実装された曲線、またはsecp224r1としても知られています。この曲線のCurveParams.Nameは「P-224」です。
=======
// Unmarshal converts a point, serialized by [Marshal], into an x, y pair. It is
// an error if the point is not in uncompressed form, is not on the curve, or is
// the point at infinity. On error, x = nil.
//
// Deprecated: for ECDH, use the crypto/ecdh package. This function accepts an
// encoding equivalent to that of the NewPublicKey methods in crypto/ecdh.
func Unmarshal(curve Curve, data []byte) (x, y *big.Int)

// UnmarshalCompressed converts a point, serialized by [MarshalCompressed], into
// an x, y pair. It is an error if the point is not in compressed form, is not
// on the curve, or is the point at infinity. On error, x = nil.
func UnmarshalCompressed(curve Curve, data []byte) (x, y *big.Int)

// P224 returns a [Curve] which implements NIST P-224 (FIPS 186-3, section D.2.2),
// also known as secp224r1. The CurveParams.Name of this [Curve] is "P-224".
>>>>>>> upstream/master
//
// この関数の複数の呼び出しは同じ値を返すため、等価性のチェックやスイッチ文に使用することができます。
//
// 暗号操作は一定の時間で実装されています。
func P224() Curve

<<<<<<< HEAD
// P256は、NIST P-256（FIPS 186-3、セクション D.2.3）またはsecp256r1またはprime256v1としても知られる、
// "P-256"という名前のCurveParams.Nameを実装したCurveを返します。
=======
// P256 returns a [Curve] which implements NIST P-256 (FIPS 186-3, section D.2.3),
// also known as secp256r1 or prime256v1. The CurveParams.Name of this [Curve] is
// "P-256".
>>>>>>> upstream/master
//
// この関数を複数回呼び出しても同じ値が返されるため、等値チェックやswitch文で使用することができます。
//
// 暗号操作は定数時間アルゴリズムを使用して実装されています。
func P256() Curve

<<<<<<< HEAD
// P384はNIST P-384 (FIPS 186-3、セクションD.2.4)、別名secp384r1を実装するCurveを返します。このCurveのCurveParams.Nameは "P-384" です。
=======
// P384 returns a [Curve] which implements NIST P-384 (FIPS 186-3, section D.2.4),
// also known as secp384r1. The CurveParams.Name of this [Curve] is "P-384".
>>>>>>> upstream/master
//
// この関数の複数の呼び出しは同じ値を返すため、等しさのチェックやスイッチ文に使用できます。
//
// 暗号操作は一定時間アルゴリズムを使用して実装されています。
func P384() Curve

<<<<<<< HEAD
// P521は、NIST P-521（FIPS 186-3、セクションD.2.5）またはsecp521r1としても知られるカーブを返します。
// このカーブのCurveParams.Nameは「P-521」です。
=======
// P521 returns a [Curve] which implements NIST P-521 (FIPS 186-3, section D.2.5),
// also known as secp521r1. The CurveParams.Name of this [Curve] is "P-521".
>>>>>>> upstream/master
//
// この関数を複数回呼び出しても同じ値が返されるため、等価性のチェックやswitch文に使用できます。
//
// 暗号操作は一定時間アルゴリズムを使用して実装されています。
func P521() Curve

// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements encoding/decoding of Floats.

package big

<<<<<<< HEAD
// GobEncode implements the [encoding/gob.GobEncoder] interface.
// The [Float] value and all its attributes (precision,
// rounding mode, accuracy) are marshaled.
func (x *Float) GobEncode() ([]byte, error)

// GobDecode implements the [encoding/gob.GobDecoder] interface.
// The result is rounded per the precision and rounding mode of
// z unless z's precision is 0, in which case z is set exactly
// to the decoded value.
func (z *Float) GobDecode(buf []byte) error

// MarshalText implements the [encoding.TextMarshaler] interface.
// Only the [Float] value is marshaled (in full precision), other
// attributes such as precision or accuracy are ignored.
func (x *Float) MarshalText() (text []byte, err error)

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
// The result is rounded per the precision and rounding mode of z.
// If z's precision is 0, it is changed to 64 before rounding takes
// effect.
=======
// GobEncodeは、gob.GobEncoderインターフェースを実装します。
// Floatの値とそのすべての属性（精度、丸めモード、精度）がマーシャルされます。
func (x *Float) GobEncode() ([]byte, error)

// GobDecodeは、gob.GobDecoderインターフェースを実装します。
// 結果は、zの精度と丸めモードに従って丸められます。
// ただし、zの精度が0の場合、zは正確にデコードされた値に設定されます。
func (z *Float) GobDecode(buf []byte) error

// MarshalTextは、encoding.TextMarshalerインターフェースを実装します。
// Floatの値のみが（全精度で）マーシャルされ、精度や精度などの他の属性は無視されます。
func (x *Float) MarshalText() (text []byte, err error)

// UnmarshalTextは、encoding.TextUnmarshalerインターフェースを実装します。
// 結果は、zの精度と丸めモードに従って丸められます。
// ただし、zの精度が0の場合、丸めが効く前に64に変更されます。
>>>>>>> release-branch.go1.21
func (z *Float) UnmarshalText(text []byte) error

// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements encoding/decoding of Floats.

package big

// GobEncodeは、[encoding/gob.GobEncoder] インターフェースを実装します。
// [Float] の値とそのすべての属性（精度、丸めモード、精度）がマーシャルされます。
func (x *Float) GobEncode() ([]byte, error)

// GobDecodeは、[encoding/gob.GobDecoder] インターフェースを実装します。
// 結果は、zの精度と丸めモードに従って丸められます。
// ただし、zの精度が0の場合、zは正確にデコードされた値に設定されます。
func (z *Float) GobDecode(buf []byte) error

<<<<<<< HEAD
// MarshalTextは、[encoding.TextMarshaler] インターフェースを実装します。
// [Float] の値のみが（全精度で）マーシャルされ、精度や精度などの他の属性は無視されます。
=======
// AppendText implements the [encoding.TextAppender] interface.
// Only the [Float] value is marshaled (in full precision), other
// attributes such as precision or accuracy are ignored.
func (x *Float) AppendText(b []byte) ([]byte, error)

// MarshalText implements the [encoding.TextMarshaler] interface.
// Only the [Float] value is marshaled (in full precision), other
// attributes such as precision or accuracy are ignored.
>>>>>>> upstream/release-branch.go1.25
func (x *Float) MarshalText() (text []byte, err error)

// UnmarshalTextは、[encoding.TextUnmarshaler] インターフェースを実装します。
// 結果は、zの精度と丸めモードに従って丸められます。
// ただし、zの精度が0の場合、丸めが効く前に64に変更されます。
func (z *Float) UnmarshalText(text []byte) error

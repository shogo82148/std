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

// AppendTextは [encoding.TextAppender] インターフェースを実装します。
// [Float] 値のみがマーシャルされ（完全な精度で）、
// 精度や正確性などの他の属性は無視されます。
func (x *Float) AppendText(b []byte) ([]byte, error)

// MarshalTextは [encoding.TextMarshaler] インターフェースを実装します。
// [Float] 値のみがマーシャルされ（完全な精度で）、
// 精度や正確性などの他の属性は無視されます。
func (x *Float) MarshalText() (text []byte, err error)

// UnmarshalTextは、[encoding.TextUnmarshaler] インターフェースを実装します。
// 結果は、zの精度と丸めモードに従って丸められます。
// ただし、zの精度が0の場合、丸めが効く前に64に変更されます。
func (z *Float) UnmarshalText(text []byte) error

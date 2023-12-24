// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements encoding/decoding of Ints.

package big

<<<<<<< HEAD
// GobEncode implements the [encoding/gob.GobEncoder] interface.
func (x *Int) GobEncode() ([]byte, error)

// GobDecode implements the [encoding/gob.GobDecoder] interface.
func (z *Int) GobDecode(buf []byte) error

// MarshalText implements the [encoding.TextMarshaler] interface.
func (x *Int) MarshalText() (text []byte, err error)

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (z *Int) UnmarshalText(text []byte) error

// MarshalJSON implements the [encoding/json.Marshaler] interface.
func (x *Int) MarshalJSON() ([]byte, error)

// UnmarshalJSON implements the [encoding/json.Unmarshaler] interface.
=======
// GobEncodeは、gob.GobEncoderインターフェースを実装します。
func (x *Int) GobEncode() ([]byte, error)

// GobDecodeは、gob.GobDecoderインターフェースを実装します。
func (z *Int) GobDecode(buf []byte) error

// MarshalTextは、encoding.TextMarshalerインターフェースを実装します。
func (x *Int) MarshalText() (text []byte, err error)

// UnmarshalTextは、encoding.TextUnmarshalerインターフェースを実装します。
func (z *Int) UnmarshalText(text []byte) error

// MarshalJSONは、json.Marshalerインターフェースを実装します。
func (x *Int) MarshalJSON() ([]byte, error)

// UnmarshalJSONは、json.Unmarshalerインターフェースを実装します。
>>>>>>> release-branch.go1.21
func (z *Int) UnmarshalJSON(text []byte) error

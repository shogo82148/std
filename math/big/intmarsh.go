// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements encoding/decoding of Ints.

package big

// GobEncodeは、[encoding/gob.GobEncoder] インターフェースを実装します。
func (x *Int) GobEncode() ([]byte, error)

// GobDecodeは、[encoding/gob.GobDecoder] インターフェースを実装します。
func (z *Int) GobDecode(buf []byte) error

// MarshalTextは、[encoding.TextMarshaler] インターフェースを実装します。
func (x *Int) MarshalText() (text []byte, err error)

// UnmarshalTextは、[encoding.TextUnmarshaler] インターフェースを実装します。
func (z *Int) UnmarshalText(text []byte) error

// MarshalJSONは、[json.Marshaler] インターフェースを実装します。
func (x *Int) MarshalJSON() ([]byte, error)

// UnmarshalJSONは、[json.Unmarshaler] インターフェースを実装します。
func (z *Int) UnmarshalJSON(text []byte) error

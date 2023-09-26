// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements encoding/decoding of Ints.

package big

// Gob codec version. Permits backward-compatible changes to the encoding.

// GobEncode implements the gob.GobEncoder interface.
func (x *Int) GobEncode() ([]byte, error)

// GobDecode implements the gob.GobDecoder interface.
func (z *Int) GobDecode(buf []byte) error

// MarshalText implements the encoding.TextMarshaler interface.
func (x *Int) MarshalText() (text []byte, err error)

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (z *Int) UnmarshalText(text []byte) error

// MarshalJSON implements the json.Marshaler interface.
func (x *Int) MarshalJSON() ([]byte, error)

// UnmarshalJSON implements the json.Unmarshaler interface.
func (z *Int) UnmarshalJSON(text []byte) error

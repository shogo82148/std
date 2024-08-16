// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements encoding/decoding of Ints.

package big

// GobEncode implements the [encoding/gob.GobEncoder] interface.
func (x *Int) GobEncode() ([]byte, error)

// GobDecode implements the [encoding/gob.GobDecoder] interface.
func (z *Int) GobDecode(buf []byte) error

// AppendText implements the [encoding.TextAppender] interface.
func (x *Int) AppendText(b []byte) (text []byte, err error)

// MarshalText implements the [encoding.TextMarshaler] interface.
func (x *Int) MarshalText() (text []byte, err error)

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (z *Int) UnmarshalText(text []byte) error

// MarshalJSON implements the [encoding/json.Marshaler] interface.
func (x *Int) MarshalJSON() ([]byte, error)

// UnmarshalJSON implements the [encoding/json.Unmarshaler] interface.
func (z *Int) UnmarshalJSON(text []byte) error

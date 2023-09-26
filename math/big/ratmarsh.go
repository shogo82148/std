// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements encoding/decoding of Rats.

package big

// Gob codec version. Permits backward-compatible changes to the encoding.

// GobEncode implements the gob.GobEncoder interface.
func (x *Rat) GobEncode() ([]byte, error)

// GobDecode implements the gob.GobDecoder interface.
func (z *Rat) GobDecode(buf []byte) error

// MarshalText implements the encoding.TextMarshaler interface.
func (x *Rat) MarshalText() (text []byte, err error)

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (z *Rat) UnmarshalText(text []byte) error

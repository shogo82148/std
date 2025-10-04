// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements encoding/decoding of Rats.

package big

// GobEncodeは、[encoding/gob.GobEncoder] インターフェースを実装します。
func (x *Rat) GobEncode() ([]byte, error)

// GobDecodeは、[encoding/gob.GobDecoder] インターフェースを実装します。
func (z *Rat) GobDecode(buf []byte) error

<<<<<<< HEAD
// MarshalTextは、[encoding.TextMarshaler] インターフェースを実装します。
=======
// AppendText implements the [encoding.TextAppender] interface.
func (x *Rat) AppendText(b []byte) ([]byte, error)

// MarshalText implements the [encoding.TextMarshaler] interface.
>>>>>>> upstream/release-branch.go1.25
func (x *Rat) MarshalText() (text []byte, err error)

// UnmarshalTextは、[encoding.TextUnmarshaler] インターフェースを実装します。
func (z *Rat) UnmarshalText(text []byte) error

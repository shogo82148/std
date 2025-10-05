// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json

import (
	"github.com/shogo82148/std/encoding/json/jsontext"
)

// Marshaler is implemented by types that can marshal themselves.
// It is recommended that types implement [MarshalerTo] unless the implementation
// is trying to avoid a hard dependency on the "jsontext" package.
//
// It is recommended that implementations return a buffer that is safe
// for the caller to retain and potentially mutate.
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// MarshalerTo is implemented by types that can marshal themselves.
// It is recommended that types implement MarshalerTo instead of [Marshaler]
// since this is both more performant and flexible.
// If a type implements both Marshaler and MarshalerTo,
// then MarshalerTo takes precedence. In such a case, both implementations
// should aim to have equivalent behavior for the default marshal options.
//
// The implementation must write only one JSON value to the Encoder and
// must not retain the pointer to [jsontext.Encoder].
type MarshalerTo interface {
	MarshalJSONTo(*jsontext.Encoder) error
}

// Unmarshaler is implemented by types that can unmarshal themselves.
// It is recommended that types implement [UnmarshalerFrom] unless the implementation
// is trying to avoid a hard dependency on the "jsontext" package.
//
// The input can be assumed to be a valid encoding of a JSON value
// if called from unmarshal functionality in this package.
// UnmarshalJSON must copy the JSON data if it is retained after returning.
// It is recommended that UnmarshalJSON implement merge semantics when
// unmarshaling into a pre-populated value.
//
// Implementations must not retain or mutate the input []byte.
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

// UnmarshalerFrom is implemented by types that can unmarshal themselves.
// It is recommended that types implement UnmarshalerFrom instead of [Unmarshaler]
// since this is both more performant and flexible.
// If a type implements both Unmarshaler and UnmarshalerFrom,
// then UnmarshalerFrom takes precedence. In such a case, both implementations
// should aim to have equivalent behavior for the default unmarshal options.
//
// The implementation must read only one JSON value from the Decoder.
// It is recommended that UnmarshalJSONFrom implement merge semantics when
// unmarshaling into a pre-populated value.
//
// Implementations must not retain the pointer to [jsontext.Decoder].
type UnmarshalerFrom interface {
	UnmarshalJSONFrom(*jsontext.Decoder) error
}

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
// is trying to avoid directly depending on the "jsontext" package.
//
// Implementations should return a buffer that is safe
// for the caller to retain and potentially mutate.
//
// If the returned error is a [SemanticError], then unpopulated fields
// of the error may be populated by [json] with additional context.
// Errors of other types are wrapped within a [SemanticError].
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// MarshalerTo is implemented by types that can marshal themselves.
// It is recommended that types implement MarshalerTo instead of [Marshaler]
// since it is both more performant and more flexible.
// If a type implements both Marshaler and MarshalerTo,
// then MarshalerTo takes precedence. In such a case, both implementations
// should aim to have equivalent behavior for the default marshal options.
//
// The implementation must write only one JSON value to the Encoder.
// Alternatively, it may return [errors.ErrUnsupported] without mutating
// the Encoder. The "json" package calling the method will
// use the next available JSON representation for the receiver type,
// as described in [Marshal].
// Implementations must not retain the pointer to [jsontext.Encoder].
//
// If the returned error is a [SemanticError], then unpopulated fields
// of the error may be populated by [json] with additional context.
// Errors of other types are wrapped within a [SemanticError],
// except for IO errors.
//
// The MarshalJSONTo method should not be called directly as it may
// return sentinel errors that need special handling.
// Users should instead call [MarshalEncode], which handles such cases.
type MarshalerTo interface {
	MarshalJSONTo(*jsontext.Encoder) error
}

// Unmarshaler is implemented by types that can unmarshal themselves.
// It is recommended that types implement [UnmarshalerFrom] unless the implementation
// is trying to avoid a direct dependency on the "jsontext" package.
//
// The input can be assumed to be a valid encoding of a JSON value
// if called from unmarshal functionality in this package.
// It is recommended that UnmarshalJSON implement merge semantics
// when unmarshaling into a pre-populated value, as described in [Unmarshal].
//
// Implementations must not retain or mutate the input []byte.
//
// If the returned error is a [SemanticError], then unpopulated fields
// of the error may be populated by [json] with additional context.
// Errors of other types are wrapped within a [SemanticError].
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

// UnmarshalerFrom is implemented by types that can unmarshal themselves.
// It is recommended that types implement UnmarshalerFrom instead of [Unmarshaler]
// since this is both more performant and more flexible.
// If a type implements both Unmarshaler and UnmarshalerFrom,
// then UnmarshalerFrom takes precedence. In such a case, both implementations
// should aim to have equivalent behavior for the default unmarshal options.
//
// The implementation must read only one JSON value from the Decoder.
// It is recommended that UnmarshalJSONFrom implement merge semantics when
// unmarshaling into a pre-populated value, as described in [Unmarshal].
// Alternatively, it may return [errors.ErrUnsupported] without mutating
// the Decoder. The "json" package calling the method will
// use the next available JSON representation for the receiver type.
// Implementations must not retain the pointer to [jsontext.Decoder].
//
// If the returned error is a [SemanticError], then unpopulated fields
// of the error may be populated by [json] with additional context.
// Errors of other types are wrapped within a [SemanticError],
// except for [jsontext.SyntacticError]s and IO errors.
//
// The UnmarshalJSONFrom method should not be called directly as it may
// return sentinel errors that need special handling.
// Users should instead call [UnmarshalDecode], which handles such cases.
type UnmarshalerFrom interface {
	UnmarshalJSONFrom(*jsontext.Decoder) error
}

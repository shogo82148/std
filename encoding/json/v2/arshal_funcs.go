// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json

import (
	"github.com/shogo82148/std/encoding/json/jsontext"
)

// Marshalers is a list of functions that may override the marshal behavior
// of specific types. Populate [WithMarshalers] to use it with
// [Marshal], [MarshalWrite], or [MarshalEncode].
// A nil *Marshalers is equivalent to an empty list.
// There are no exported fields or methods on Marshalers.
type Marshalers = typedMarshalers

// JoinMarshalers constructs a flattened list of marshal functions.
// If multiple functions in the list are applicable for a value of a given type,
// then those earlier in the list take precedence over those that come later.
// If a function returns [errors.ErrUnsupported],
// then the next applicable function is called,
// otherwise the default marshaling behavior is used.
//
// For example:
//
//	m1 := JoinMarshalers(f1, f2)
//	m2 := JoinMarshalers(f0, m1, f3)     // equivalent to m3
//	m3 := JoinMarshalers(f0, f1, f2, f3) // equivalent to m2
func JoinMarshalers(ms ...*Marshalers) *Marshalers

// Unmarshalers is a list of functions that may override the unmarshal behavior
// of specific types. Populate [WithUnmarshalers] to use it with
// [Unmarshal], [UnmarshalRead], or [UnmarshalDecode].
// A nil *Unmarshalers is equivalent to an empty list.
// There are no exported fields or methods on Unmarshalers.
type Unmarshalers = typedUnmarshalers

// JoinUnmarshalers constructs a flattened list of unmarshal functions.
// If multiple functions in the list are applicable for a value of a given type,
// then those earlier in the list take precedence over those that come later.
// If a function returns [errors.ErrUnsupported],
// then the next applicable function is called,
// otherwise the default unmarshaling behavior is used.
//
// For example:
//
//	u1 := JoinUnmarshalers(f1, f2)
//	u2 := JoinUnmarshalers(f0, u1, f3)     // equivalent to u3
//	u3 := JoinUnmarshalers(f0, f1, f2, f3) // equivalent to u2
func JoinUnmarshalers(us ...*Unmarshalers) *Unmarshalers

// MarshalFunc constructs a type-specific marshaler that
// specifies how to marshal values of type T.
// T can be any type except a named pointer.
// The function is always provided with a non-nil pointer value
// if T is an interface or pointer type.
//
// Implementations must follow the requirements of [Marshaler].
//
// Implementations must not retain the value of T.
func MarshalFunc[T any](fn func(T) ([]byte, error)) *Marshalers

// MarshalToFunc constructs a type-specific marshaler that
// specifies how to marshal values of type T.
// T can be any type except a named pointer.
// The function is always provided with a non-nil pointer value
// if T is an interface or pointer type.
//
// Implementations must follow the requirements of [MarshalerTo].
//
// Implementations must not retain the pointer to [jsontext.Encoder] or the
// value of T.
func MarshalToFunc[T any](fn func(*jsontext.Encoder, T) error) *Marshalers

// UnmarshalFunc constructs a type-specific unmarshaler that
// specifies how to unmarshal values of type T.
// T must be an unnamed pointer or an interface type.
// The function is always provided with a non-nil pointer value.
//
// Implementations must follow the requirements of [Unmarshaler].
//
// Implementations must not retain the value of T.
func UnmarshalFunc[T any](fn func([]byte, T) error) *Unmarshalers

// UnmarshalFromFunc constructs a type-specific unmarshaler that
// specifies how to unmarshal values of type T.
// T must be an unnamed pointer or an interface type.
// The function is always provided with a non-nil pointer value.
//
// Implementations must follow the requirements of [UnmarshalerFrom].
//
// Implementations must not retain the pointer to [jsontext.Decoder] or the
// value of T.
func UnmarshalFromFunc[T any](fn func(*jsontext.Decoder, T) error) *Unmarshalers

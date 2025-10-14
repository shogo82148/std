// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/reflect"

	"github.com/shogo82148/std/encoding/json/jsontext"
)

// ErrUnknownName indicates that a JSON object member could not be
// unmarshaled because the name is not known to the target Go struct.
// This error is directly wrapped within a [SemanticError] when produced.
//
// The name of an unknown JSON object member can be extracted as:
//
//	err := ...
//	if serr, ok := errors.AsType[json.SemanticError](err); ok && serr.Err == json.ErrUnknownName {
//		ptr := serr.JSONPointer // JSON pointer to unknown name
//		name := ptr.LastToken() // unknown name itself
//		...
//	}
//
// This error is only returned if [RejectUnknownMembers] is true.
var ErrUnknownName = errors.New("unknown object member name")

// SemanticError describes an error determining the meaning
// of JSON data as Go data or vice-versa.
//
// If a [Marshaler], [MarshalerTo], [Unmarshaler], or [UnmarshalerFrom] method
// returns a SemanticError when called by the [json] package,
// then the ByteOffset, JSONPointer, and GoType fields are automatically
// populated by the calling context if they are the zero value.
//
// The contents of this error as produced by this package may change over time.
type SemanticError struct {
	requireKeyedLiterals
	nonComparable

	action string

	// ByteOffset indicates that an error occurred after this byte offset.
	ByteOffset int64
	// JSONPointer indicates that an error occurred within this JSON value
	// as indicated using the JSON Pointer notation (see RFC 6901).
	JSONPointer jsontext.Pointer

	// JSONKind is the JSON kind that could not be handled.
	JSONKind jsontext.Kind
	// JSONValue is the JSON number or string that could not be unmarshaled.
	// It is not populated during marshaling.
	JSONValue jsontext.Value
	// GoType is the Go type that could not be handled.
	GoType reflect.Type

	// Err is the underlying error.
	Err error
}

func (e *SemanticError) Error() string

func (e *SemanticError) Unwrap() error

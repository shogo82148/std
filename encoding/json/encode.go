// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package json implements encoding and decoding of JSON objects as defined in
// RFC 4627.
//
// See "JSON and Go" for an introduction to this package:
// http://golang.org/doc/articles/json_and_go.html
package json

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/reflect"
)

// Marshal returns the JSON encoding of v.
//
// Marshal traverses the value v recursively.
// If an encountered value implements the Marshaler interface
// and is not a nil pointer, Marshal calls its MarshalJSON method
// to produce JSON.  The nil pointer exception is not strictly necessary
// but mimics a similar, necessary exception in the behavior of
// UnmarshalJSON.
//
// Otherwise, Marshal uses the following type-dependent default encodings:
//
// Boolean values encode as JSON booleans.
//
// Floating point and integer values encode as JSON numbers.
//
// String values encode as JSON strings, with each invalid UTF-8 sequence
// replaced by the encoding of the Unicode replacement character U+FFFD.
// The angle brackets "<" and ">" are escaped to "\u003c" and "\u003e"
// to keep some browsers from misinterpreting JSON output as HTML.
//
// Array and slice values encode as JSON arrays, except that
// []byte encodes as a base64-encoded string, and a nil slice
// encodes as the null JSON object.
//
// Struct values encode as JSON objects. Each exported struct field
// becomes a member of the object unless
//   - the field's tag is "-", or
//   - the field is empty and its tag specifies the "omitempty" option.
//
// The empty values are false, 0, any
// nil pointer or interface value, and any array, slice, map, or string of
// length zero. The object's default key string is the struct field name
// but can be specified in the struct field's tag value. The "json" key in
// struct field's tag value is the key name, followed by an optional comma
// and options. Examples:
//
//	// Field is ignored by this package.
//	Field int `json:"-"`
//
//	// Field appears in JSON as key "myName".
//	Field int `json:"myName"`
//
//	// Field appears in JSON as key "myName" and
//	// the field is omitted from the object if its value is empty,
//	// as defined above.
//	Field int `json:"myName,omitempty"`
//
//	// Field appears in JSON as key "Field" (the default), but
//	// the field is skipped if empty.
//	// Note the leading comma.
//	Field int `json:",omitempty"`
//
// The "string" option signals that a field is stored as JSON inside a
// JSON-encoded string.  This extra level of encoding is sometimes
// used when communicating with JavaScript programs:
//
//	Int64String int64 `json:",string"`
//
// The key name will be used if it's a non-empty string consisting of
// only Unicode letters, digits, dollar signs, percent signs, hyphens,
// underscores and slashes.
//
// Map values encode as JSON objects.
// The map's key type must be string; the object keys are used directly
// as map keys.
//
// Pointer values encode as the value pointed to.
// A nil pointer encodes as the null JSON object.
//
// Interface values encode as the value contained in the interface.
// A nil interface value encodes as the null JSON object.
//
// Channel, complex, and function values cannot be encoded in JSON.
// Attempting to encode such a value causes Marshal to return
// an InvalidTypeError.
//
// JSON cannot represent cyclic data structures and Marshal does not
// handle them.  Passing cyclic structures to Marshal will result in
// an infinite recursion.
func Marshal(v interface{}) ([]byte, error)

// MarshalIndent is like Marshal but applies Indent to format the output.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

// HTMLEscape appends to dst the JSON-encoded src with <, >, and &
// characters inside string literals changed to \u003c, \u003e, \u0026
// so that the JSON will be safe to embed inside HTML <script> tags.
// For historical reasons, web browsers don't honor standard HTML
// escaping within <script> tags, so an alternative JSON encoding must
// be used.
func HTMLEscape(dst *bytes.Buffer, src []byte)

// Marshaler is the interface implemented by objects that
// can marshal themselves into valid JSON.
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string

type UnsupportedValueError struct {
	Value reflect.Value
	Str   string
}

func (e *UnsupportedValueError) Error() string

type InvalidUTF8Error struct {
	S string
}

func (e *InvalidUTF8Error) Error() string

type MarshalerError struct {
	Type reflect.Type
	Err  error
}

func (e *MarshalerError) Error() string

// An encodeState encodes JSON into a bytes.Buffer.

// stringValues is a slice of reflect.Value holding *reflect.StringValue.
// It implements the methods to sort by string.

// encodeField contains information about how to encode a field of a
// struct.

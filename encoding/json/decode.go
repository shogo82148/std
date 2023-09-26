// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Represents JSON data structure using native Go types: booleans, floats,
// strings, arrays, and maps.

package json

import (
	"github.com/shogo82148/std/reflect"
)

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v.
//
// Unmarshal uses the inverse of the encodings that
// Marshal uses, allocating maps, slices, and pointers as necessary,
// with the following additional rules:
//
// To unmarshal JSON into a pointer, Unmarshal first handles the case of
// the JSON being the JSON literal null.  In that case, Unmarshal sets
// the pointer to nil.  Otherwise, Unmarshal unmarshals the JSON into
// the value pointed at by the pointer.  If the pointer is nil, Unmarshal
// allocates a new value for it to point to.
//
// To unmarshal JSON into an interface value, Unmarshal unmarshals
// the JSON into the concrete value contained in the interface value.
// If the interface value is nil, that is, has no concrete value stored in it,
// Unmarshal stores one of these in the interface value:
//
//	bool, for JSON booleans
//	float64, for JSON numbers
//	string, for JSON strings
//	[]interface{}, for JSON arrays
//	map[string]interface{}, for JSON objects
//	nil for JSON null
//
// If a JSON value is not appropriate for a given target type,
// or if a JSON number overflows the target type, Unmarshal
// skips that field and completes the unmarshalling as best it can.
// If no more serious errors are encountered, Unmarshal returns
// an UnmarshalTypeError describing the earliest such error.
func Unmarshal(data []byte, v interface{}) error

// Unmarshaler is the interface implemented by objects
// that can unmarshal a JSON description of themselves.
// The input can be assumed to be a valid JSON object
// encoding.  UnmarshalJSON must copy the JSON data
// if it wishes to retain the data after returning.
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

// An UnmarshalTypeError describes a JSON value that was
// not appropriate for a value of a specific Go type.
type UnmarshalTypeError struct {
	Value string
	Type  reflect.Type
}

func (e *UnmarshalTypeError) Error() string

// An UnmarshalFieldError describes a JSON object key that
// led to an unexported (and therefore unwritable) struct field.
type UnmarshalFieldError struct {
	Key   string
	Type  reflect.Type
	Field reflect.StructField
}

func (e *UnmarshalFieldError) Error() string

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string

// decodeState represents the state while decoding a JSON value.

// errPhase is used for errors that should not happen unless
// there is a bug in the JSON decoder or something is editing
// the data slice while the decoder executes.

// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext

import (
	"github.com/shogo82148/std/io"
)

// Encoder is a streaming encoder from raw JSON tokens and values.
// It is used to write a stream of top-level JSON values,
// each terminated with a newline character.
//
// [Encoder.WriteToken] and [Encoder.WriteValue] calls may be interleaved.
// For example, the following JSON value:
//
//	{"name":"value","array":[null,false,true,3.14159],"object":{"k":"v"}}
//
// can be composed with the following calls (ignoring errors for brevity):
//
//	e.WriteToken(BeginObject)        // {
//	e.WriteToken(String("name"))     // "name"
//	e.WriteToken(String("value"))    // "value"
//	e.WriteValue(Value(`"array"`))   // "array"
//	e.WriteToken(BeginArray)         // [
//	e.WriteToken(Null)               // null
//	e.WriteToken(False)              // false
//	e.WriteValue(Value("true"))      // true
//	e.WriteToken(Float(3.14159))     // 3.14159
//	e.WriteToken(EndArray)           // ]
//	e.WriteValue(Value(`"object"`))  // "object"
//	e.WriteValue(Value(`{"k":"v"}`)) // {"k":"v"}
//	e.WriteToken(EndObject)          // }
//
// The above is one of many possible sequence of calls and
// may not represent the most sensible method to call for any given token/value.
// For example, it is probably more common to call [Encoder.WriteToken] with a string
// for object names.
type Encoder struct {
	s encoderState
}

// NewEncoder constructs a new streaming encoder writing to w
// configured with the provided options.
// It flushes the internal buffer when the buffer is sufficiently full or
// when a top-level value has been written.
//
// If w is a [bytes.Buffer], then the encoder appends directly into the buffer
// without copying the contents from an intermediate buffer.
func NewEncoder(w io.Writer, opts ...Options) *Encoder

// Reset resets an encoder such that it is writing afresh to w and
// configured with the provided options. Reset must not be called on
// a Encoder passed to the [encoding/json/v2.MarshalerTo.MarshalJSONTo] method
// or the [encoding/json/v2.MarshalToFunc] function.
func (e *Encoder) Reset(w io.Writer, opts ...Options)

// Options returns the options used to construct the decoder and
// may additionally contain semantic options passed to a
// [encoding/json/v2.MarshalEncode] call.
//
// If operating within
// a [encoding/json/v2.MarshalerTo.MarshalJSONTo] method call or
// a [encoding/json/v2.MarshalToFunc] function call,
// then the returned options are only valid within the call.
func (e *Encoder) Options() Options

// WriteToken writes the next token and advances the internal write offset.
//
// The provided token kind must be consistent with the JSON grammar.
// For example, it is an error to provide a number when the encoder
// is expecting an object name (which is always a string), or
// to provide an end object delimiter when the encoder is finishing an array.
// If the provided token is invalid, then it reports a [SyntacticError] and
// the internal state remains unchanged. The offset reported
// in [SyntacticError] will be relative to the [Encoder.OutputOffset].
func (e *Encoder) WriteToken(t Token) error

// WriteValue writes the next raw value and advances the internal write offset.
// The Encoder does not simply copy the provided value verbatim, but
// parses it to ensure that it is syntactically valid and reformats it
// according to how the Encoder is configured to format whitespace and strings.
// If [AllowInvalidUTF8] is specified, then any invalid UTF-8 is mangled
// as the Unicode replacement character, U+FFFD.
//
// The provided value kind must be consistent with the JSON grammar
// (see examples on [Encoder.WriteToken]). If the provided value is invalid,
// then it reports a [SyntacticError] and the internal state remains unchanged.
// The offset reported in [SyntacticError] will be relative to the
// [Encoder.OutputOffset] plus the offset into v of any encountered syntax error.
func (e *Encoder) WriteValue(v Value) error

// OutputOffset returns the current output byte offset. It gives the location
// of the next byte immediately after the most recently written token or value.
// The number of bytes actually written to the underlying [io.Writer] may be less
// than this offset due to internal buffering effects.
func (e *Encoder) OutputOffset() int64

// AvailableBuffer returns a zero-length buffer with a possible non-zero capacity.
// This buffer is intended to be used to populate a [Value]
// being passed to an immediately succeeding [Encoder.WriteValue] call.
//
// Example usage:
//
//	b := d.AvailableBuffer()
//	b = append(b, '"')
//	b = appendString(b, v) // append the string formatting of v
//	b = append(b, '"')
//	... := d.WriteValue(b)
//
// It is the user's responsibility to ensure that the value is valid JSON.
func (e *Encoder) AvailableBuffer() []byte

// StackDepth returns the depth of the state machine for written JSON data.
// Each level on the stack represents a nested JSON object or array.
// It is incremented whenever an [BeginObject] or [BeginArray] token is encountered
// and decremented whenever an [EndObject] or [EndArray] token is encountered.
// The depth is zero-indexed, where zero represents the top-level JSON value.
func (e *Encoder) StackDepth() int

// StackIndex returns information about the specified stack level.
// It must be a number between 0 and [Encoder.StackDepth], inclusive.
// For each level, it reports the kind:
//
//   - [KindInvalid] for a level of zero,
//   - [KindBeginObject] for a level representing a JSON object, and
//   - [KindBeginArray] for a level representing a JSON array.
//
// It also reports the length of that JSON object or array.
// Each name and value in a JSON object is counted separately,
// so the effective number of members would be half the length.
// A complete JSON object must have an even length.
func (e *Encoder) StackIndex(i int) (Kind, int64)

// StackPointer returns a JSON Pointer (RFC 6901) to the most recently written value.
func (e *Encoder) StackPointer() Pointer

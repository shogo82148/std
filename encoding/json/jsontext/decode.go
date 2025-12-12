// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext

import (
	"github.com/shogo82148/std/io"
)

// Decoder is a streaming decoder for raw JSON tokens and values.
// It is used to read a stream of top-level JSON values,
// each separated by optional whitespace characters.
//
// [Decoder.ReadToken] and [Decoder.ReadValue] calls may be interleaved.
// For example, the following JSON value:
//
//	{"name":"value","array":[null,false,true,3.14159],"object":{"k":"v"}}
//
// can be parsed with the following calls (ignoring errors for brevity):
//
//	d.ReadToken() // {
//	d.ReadToken() // "name"
//	d.ReadToken() // "value"
//	d.ReadValue() // "array"
//	d.ReadToken() // [
//	d.ReadToken() // null
//	d.ReadToken() // false
//	d.ReadValue() // true
//	d.ReadToken() // 3.14159
//	d.ReadToken() // ]
//	d.ReadValue() // "object"
//	d.ReadValue() // {"k":"v"}
//	d.ReadToken() // }
//
// The above is one of many possible sequence of calls and
// may not represent the most sensible method to call for any given token/value.
// For example, it is probably more common to call [Decoder.ReadToken] to obtain a
// string token for object names.
type Decoder struct {
	s decoderState
}

// NewDecoder constructs a new streaming decoder reading from r.
//
// If r is a [bytes.Buffer], then the decoder parses directly from the buffer
// without first copying the contents to an intermediate buffer.
// Additional writes to the buffer must not occur while the decoder is in use.
func NewDecoder(r io.Reader, opts ...Options) *Decoder

// Reset resets a decoder such that it is reading afresh from r and
// configured with the provided options. Reset must not be called on an
// a Decoder passed to the [encoding/json/v2.UnmarshalerFrom.UnmarshalJSONFrom] method
// or the [encoding/json/v2.UnmarshalFromFunc] function.
func (d *Decoder) Reset(r io.Reader, opts ...Options)

// Options returns the options used to construct the encoder and
// may additionally contain semantic options passed to a
// [encoding/json/v2.UnmarshalDecode] call.
//
// If operating within
// a [encoding/json/v2.UnmarshalerFrom.UnmarshalJSONFrom] method call or
// a [encoding/json/v2.UnmarshalFromFunc] function call,
// then the returned options are only valid within the call.
func (d *Decoder) Options() Options

// PeekKind retrieves the next token kind, but does not advance the read offset.
//
// It returns [KindInvalid] if an error occurs. Any such error is cached until
// the next read call and it is the caller's responsibility to eventually
// follow up a PeekKind call with a read call.
func (d *Decoder) PeekKind() Kind

// SkipValue is semantically equivalent to calling [Decoder.ReadValue] and discarding
// the result except that memory is not wasted trying to hold the entire result.
func (d *Decoder) SkipValue() error

// ReadToken reads the next [Token], advancing the read offset.
// The returned token is only valid until the next Peek, Read, or Skip call.
// It returns [io.EOF] if there are no more tokens.
func (d *Decoder) ReadToken() (Token, error)

// ReadValue returns the next raw JSON value, advancing the read offset.
// The value is stripped of any leading or trailing whitespace and
// contains the exact bytes of the input, which may contain invalid UTF-8
// if [AllowInvalidUTF8] is specified.
//
// The returned value is only valid until the next Peek, Read, or Skip call and
// may not be mutated while the Decoder remains in use.
// If the decoder is currently at the end token for an object or array,
// then it reports a [SyntacticError] and the internal state remains unchanged.
// It returns [io.EOF] if there are no more values.
func (d *Decoder) ReadValue() (Value, error)

// InputOffset returns the current input byte offset. It gives the location
// of the next byte immediately after the most recently returned token or value.
// The number of bytes actually read from the underlying [io.Reader] may be more
// than this offset due to internal buffering effects.
func (d *Decoder) InputOffset() int64

// UnreadBuffer returns the data remaining in the unread buffer,
// which may contain zero or more bytes.
// The returned buffer must not be mutated while Decoder continues to be used.
// The buffer contents are valid until the next Peek, Read, or Skip call.
func (d *Decoder) UnreadBuffer() []byte

// StackDepth returns the depth of the state machine for read JSON data.
// Each level on the stack represents a nested JSON object or array.
// It is incremented whenever an [BeginObject] or [BeginArray] token is encountered
// and decremented whenever an [EndObject] or [EndArray] token is encountered.
// The depth is zero-indexed, where zero represents the top-level JSON value.
func (d *Decoder) StackDepth() int

// StackIndex returns information about the specified stack level.
// It must be a number between 0 and [Decoder.StackDepth], inclusive.
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
func (d *Decoder) StackIndex(i int) (Kind, int64)

// StackPointer returns a JSON Pointer (RFC 6901) to the most recently read value.
func (d *Decoder) StackPointer() Pointer

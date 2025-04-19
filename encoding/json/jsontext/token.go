// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext

// Token represents a lexical JSON token, which may be one of the following:
//   - a JSON literal (i.e., null, true, or false)
//   - a JSON string (e.g., "hello, world!")
//   - a JSON number (e.g., 123.456)
//   - a start or end delimiter for a JSON object (i.e., { or } )
//   - a start or end delimiter for a JSON array (i.e., [ or ] )
//
// A Token cannot represent entire array or object values, while a [Value] can.
// There is no Token to represent commas and colons since
// these structural tokens can be inferred from the surrounding context.
type Token struct {
	nonComparable

	// raw contains a reference to the raw decode buffer.
	// If non-nil, then its value takes precedence over str and num.
	// It is only valid if num == raw.previousOffsetStart().
	raw *decodeBuffer

	// str is the unescaped JSON string if num is zero.
	// Otherwise, it is "f", "i", or "u" if num should be interpreted
	// as a float64, int64, or uint64, respectively.
	str string

	// num is a float64, int64, or uint64 stored as a uint64 value.
	// It is non-zero for any JSON number in the "exact" form.
	num uint64
}

var (
	Null  Token = rawToken("null")
	False Token = rawToken("false")
	True  Token = rawToken("true")

	BeginObject Token = rawToken("{")
	EndObject   Token = rawToken("}")
	BeginArray  Token = rawToken("[")
	EndArray    Token = rawToken("]")
)

// Bool constructs a Token representing a JSON boolean.
func Bool(b bool) Token

// String constructs a Token representing a JSON string.
// The provided string should contain valid UTF-8, otherwise invalid characters
// may be mangled as the Unicode replacement character.
func String(s string) Token

// Float constructs a Token representing a JSON number.
// The values NaN, +Inf, and -Inf will be represented
// as a JSON string with the values "NaN", "Infinity", and "-Infinity".
func Float(n float64) Token

// Int constructs a Token representing a JSON number from an int64.
func Int(n int64) Token

// Uint constructs a Token representing a JSON number from a uint64.
func Uint(n uint64) Token

// Clone makes a copy of the Token such that its value remains valid
// even after a subsequent [Decoder.Read] call.
func (t Token) Clone() Token

// Bool returns the value for a JSON boolean.
// It panics if the token kind is not a JSON boolean.
func (t Token) Bool() bool

// String returns the unescaped string value for a JSON string.
// For other JSON kinds, this returns the raw JSON representation.
func (t Token) String() string

// Float returns the floating-point value for a JSON number.
// It returns a NaN, +Inf, or -Inf value for any JSON string
// with the values "NaN", "Infinity", or "-Infinity".
// It panics for all other cases.
func (t Token) Float() float64

// Int returns the signed integer value for a JSON number.
// The fractional component of any number is ignored (truncation toward zero).
// Any number beyond the representation of an int64 will be saturated
// to the closest representable value.
// It panics if the token kind is not a JSON number.
func (t Token) Int() int64

// Uint returns the unsigned integer value for a JSON number.
// The fractional component of any number is ignored (truncation toward zero).
// Any number beyond the representation of an uint64 will be saturated
// to the closest representable value.
// It panics if the token kind is not a JSON number.
func (t Token) Uint() uint64

// Kind returns the token kind.
func (t Token) Kind() Kind

// Kind represents each possible JSON token kind with a single byte,
// which is conveniently the first byte of that kind's grammar
// with the restriction that numbers always be represented with '0':
//
//   - 'n': null
//   - 'f': false
//   - 't': true
//   - '"': string
//   - '0': number
//   - '{': object start
//   - '}': object end
//   - '[': array start
//   - ']': array end
//
// An invalid kind is usually represented using 0,
// but may be non-zero due to invalid JSON data.
type Kind byte

// String prints the kind in a humanly readable fashion.
func (k Kind) String() string

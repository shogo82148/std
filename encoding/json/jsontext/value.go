// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext

// AppendFormat formats the JSON value in src and appends it to dst
// according to the specified options.
// See [Value.Format] for more details about the formatting behavior.
//
// The dst and src may overlap.
// If an error is reported, then the entirety of src is appended to dst.
func AppendFormat(dst, src []byte, opts ...Options) ([]byte, error)

// Value represents a single raw JSON value, which may be one of the following:
//   - a JSON literal (i.e., null, true, or false)
//   - a JSON string (e.g., "hello, world!")
//   - a JSON number (e.g., 123.456)
//   - an entire JSON object (e.g., {"fizz":"buzz"} )
//   - an entire JSON array (e.g., [1,2,3] )
//
// Value can represent entire array or object values, while [Token] cannot.
// Value may contain leading and/or trailing whitespace.
type Value []byte

// Clone returns a copy of v.
func (v Value) Clone() Value

// String returns the string formatting of v.
func (v Value) String() string

// IsValid reports whether the raw JSON value is syntactically valid
// according to the specified options.
//
// By default (if no options are specified), it validates according to RFC 7493.
// It verifies whether the input is properly encoded as UTF-8,
// that escape sequences within strings decode to valid Unicode codepoints, and
// that all names in each object are unique.
// It does not verify whether numbers are representable within the limits
// of any common numeric type (e.g., float64, int64, or uint64).
//
// Relevant options include:
//   - [AllowDuplicateNames]
//   - [AllowInvalidUTF8]
//
// All other options are ignored.
func (v Value) IsValid(opts ...Options) bool

// Format formats the raw JSON value in place.
//
// By default (if no options are specified), it validates according to RFC 7493
// and produces the minimal JSON representation, where
// all whitespace is elided and JSON strings use the shortest encoding.
//
// Relevant options include:
//   - [AllowDuplicateNames]
//   - [AllowInvalidUTF8]
//   - [EscapeForHTML]
//   - [EscapeForJS]
//   - [PreserveRawStrings]
//   - [CanonicalizeRawInts]
//   - [CanonicalizeRawFloats]
//   - [ReorderRawObjects]
//   - [SpaceAfterColon]
//   - [SpaceAfterComma]
//   - [Multiline]
//   - [WithIndent]
//   - [WithIndentPrefix]
//
// All other options are ignored.
//
// It is guaranteed to succeed if the value is valid according to the same options.
// If the value is already formatted, then the buffer is not mutated.
func (v *Value) Format(opts ...Options) error

// Compact removes all whitespace from the raw JSON value.
//
// It does not reformat JSON strings or numbers to use any other representation.
// To maximize the set of JSON values that can be formatted,
// this permits values with duplicate names and invalid UTF-8.
//
// Compact is equivalent to calling [Value.Format] with the following options:
//   - [AllowDuplicateNames](true)
//   - [AllowInvalidUTF8](true)
//   - [PreserveRawStrings](true)
//
// Any options specified by the caller are applied after the initial set
// and may deliberately override prior options.
func (v *Value) Compact(opts ...Options) error

// Indent reformats the whitespace in the raw JSON value so that each element
// in a JSON object or array begins on a indented line according to the nesting.
//
// It does not reformat JSON strings or numbers to use any other representation.
// To maximize the set of JSON values that can be formatted,
// this permits values with duplicate names and invalid UTF-8.
//
// Indent is equivalent to calling [Value.Format] with the following options:
//   - [AllowDuplicateNames](true)
//   - [AllowInvalidUTF8](true)
//   - [PreserveRawStrings](true)
//   - [Multiline](true)
//
// Any options specified by the caller are applied after the initial set
// and may deliberately override prior options.
func (v *Value) Indent(opts ...Options) error

// Canonicalize canonicalizes the raw JSON value according to the
// JSON Canonicalization Scheme (JCS) as defined by RFC 8785
// where it produces a stable representation of a JSON value.
//
// JSON strings are formatted to use their minimal representation,
// JSON numbers are formatted as double precision numbers according
// to some stable serialization algorithm.
// JSON object members are sorted in ascending order by name.
// All whitespace is removed.
//
// The output stability is dependent on the stability of the application data
// (see RFC 8785, Appendix E). It cannot produce stable output from
// fundamentally unstable input. For example, if the JSON value
// contains ephemeral data (e.g., a frequently changing timestamp),
// then the value is still unstable regardless of whether this is called.
//
// Canonicalize is equivalent to calling [Value.Format] with the following options:
//   - [CanonicalizeRawInts](true)
//   - [CanonicalizeRawFloats](true)
//   - [ReorderRawObjects](true)
//
// Any options specified by the caller are applied after the initial set
// and may deliberately override prior options.
//
// Note that JCS treats all JSON numbers as IEEE 754 double precision numbers.
// Any numbers with precision beyond what is representable by that form
// will lose their precision when canonicalized. For example, integer values
// beyond ±2⁵³ will lose their precision. To preserve the original representation
// of JSON integers, additionally set [CanonicalizeRawInts] to false:
//
//	v.Canonicalize(jsontext.CanonicalizeRawInts(false))
func (v *Value) Canonicalize(opts ...Options) error

// MarshalJSON returns v as the JSON encoding of v.
// It returns the stored value as the raw JSON output without any validation.
// If v is nil, then this returns a JSON null.
func (v Value) MarshalJSON() ([]byte, error)

// UnmarshalJSON sets v as the JSON encoding of b.
// It stores a copy of the provided raw JSON input without any validation.
func (v *Value) UnmarshalJSON(b []byte) error

// Kind returns the starting token kind.
// For a valid value, this will never include [KindEndObject] or [KindEndArray].
func (v Value) Kind() Kind

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsonwire

type ValueFlags uint

const (
	_ ValueFlags = (1 << iota) / 2
)

func (f *ValueFlags) Join(f2 ValueFlags)
func (f ValueFlags) IsVerbatim() bool
func (f ValueFlags) IsCanonical() bool

// ConsumeWhitespace consumes leading JSON whitespace per RFC 7159, section 2.
func ConsumeWhitespace(b []byte) (n int)

// ConsumeNull consumes the next JSON null literal per RFC 7159, section 3.
// It returns 0 if it is invalid, in which case consumeLiteral should be used.
func ConsumeNull(b []byte) int

// ConsumeFalse consumes the next JSON false literal per RFC 7159, section 3.
// It returns 0 if it is invalid, in which case consumeLiteral should be used.
func ConsumeFalse(b []byte) int

// ConsumeTrue consumes the next JSON true literal per RFC 7159, section 3.
// It returns 0 if it is invalid, in which case consumeLiteral should be used.
func ConsumeTrue(b []byte) int

// ConsumeLiteral consumes the next JSON literal per RFC 7159, section 3.
// If the input appears truncated, it returns io.ErrUnexpectedEOF.
func ConsumeLiteral(b []byte, lit string) (n int, err error)

// ConsumeSimpleString consumes the next JSON string per RFC 7159, section 7
// but is limited to the grammar for an ASCII string without escape sequences.
// It returns 0 if it is invalid or more complicated than a simple string,
// in which case consumeString should be called.
//
// It rejects '<', '>', and '&' for compatibility reasons since these were
// always escaped in the v1 implementation. Thus, if this function reports
// non-zero then we know that the string would be encoded the same way
// under both v1 or v2 escape semantics.
func ConsumeSimpleString(b []byte) (n int)

// ConsumeString consumes the next JSON string per RFC 7159, section 7.
// If validateUTF8 is false, then this allows the presence of invalid UTF-8
// characters within the string itself.
// It reports the number of bytes consumed and whether an error was encountered.
// If the input appears truncated, it returns io.ErrUnexpectedEOF.
func ConsumeString(flags *ValueFlags, b []byte, validateUTF8 bool) (n int, err error)

// ConsumeStringResumable is identical to consumeString but supports resuming
// from a previous call that returned io.ErrUnexpectedEOF.
func ConsumeStringResumable(flags *ValueFlags, b []byte, resumeOffset int, validateUTF8 bool) (n int, err error)

// AppendUnquote appends the unescaped form of a JSON string in src to dst.
// Any invalid UTF-8 within the string will be replaced with utf8.RuneError,
// but the error will be specified as having encountered such an error.
// The input must be an entire JSON string with no surrounding whitespace.
func AppendUnquote[Bytes ~[]byte | ~string](dst []byte, src Bytes) (v []byte, err error)

// UnquoteMayCopy returns the unescaped form of b.
// If there are no escaped characters, the output is simply a subslice of
// the input with the surrounding quotes removed.
// Otherwise, a new buffer is allocated for the output.
// It assumes the input is valid.
func UnquoteMayCopy(b []byte, isVerbatim bool) []byte

// ConsumeSimpleNumber consumes the next JSON number per RFC 7159, section 6
// but is limited to the grammar for a positive integer.
// It returns 0 if it is invalid or more complicated than a simple integer,
// in which case consumeNumber should be called.
func ConsumeSimpleNumber(b []byte) (n int)

type ConsumeNumberState uint

// ConsumeNumber consumes the next JSON number per RFC 7159, section 6.
// It reports the number of bytes consumed and whether an error was encountered.
// If the input appears truncated, it returns io.ErrUnexpectedEOF.
//
// Note that JSON numbers are not self-terminating.
// If the entire input is consumed, then the caller needs to consider whether
// there may be subsequent unread data that may still be part of this number.
func ConsumeNumber(b []byte) (n int, err error)

// ConsumeNumberResumable is identical to consumeNumber but supports resuming
// from a previous call that returned io.ErrUnexpectedEOF.
func ConsumeNumberResumable(b []byte, resumeOffset int, state ConsumeNumberState) (n int, _ ConsumeNumberState, err error)

// ParseUint parses b as a decimal unsigned integer according to
// a strict subset of the JSON number grammar, returning the value if valid.
// It returns (0, false) if there is a syntax error and
// returns (math.MaxUint64, false) if there is an overflow.
func ParseUint(b []byte) (v uint64, ok bool)

// ParseFloat parses a floating point number according to the Go float grammar.
// Note that the JSON number grammar is a strict subset.
//
// If the number overflows the finite representation of a float,
// then we return MaxFloat since any finite value will always be infinitely
// more accurate at representing another finite value than an infinite value.
func ParseFloat(b []byte, bits int) (v float64, ok bool)

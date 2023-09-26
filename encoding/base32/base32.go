// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package base32 implements base32 encoding as specified by RFC 4648.
package base32

import (
	"github.com/shogo82148/std/io"
)

// An Encoding is a radix 32 encoding/decoding scheme, defined by a
// 32-character alphabet. The most common is the "base32" encoding
// introduced for SASL GSSAPI and standardized in RFC 4648.
// The alternate "base32hex" encoding is used in DNSSEC.
type Encoding struct {
	encode    [32]byte
	decodeMap [256]uint8
	padChar   rune
}

const (
	StdPadding rune = '='
	NoPadding  rune = -1
)

// NewEncoding returns a new padded Encoding defined by the given alphabet,
// which must be a 32-byte string that contains unique byte values and
// does not contain the padding character or CR / LF ('\r', '\n').
// The alphabet is treated as a sequence of byte values
// without any special treatment for multi-byte UTF-8.
// The resulting Encoding uses the default padding character ('='),
// which may be changed or disabled via [Encoding.WithPadding].
func NewEncoding(encoder string) *Encoding

// StdEncoding is the standard base32 encoding, as defined in RFC 4648.
var StdEncoding = NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567")

// HexEncoding is the “Extended Hex Alphabet” defined in RFC 4648.
// It is typically used in DNS.
var HexEncoding = NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUV")

// WithPadding creates a new encoding identical to enc except
// with a specified padding character, or NoPadding to disable padding.
// The padding character must not be '\r' or '\n',
// must not be contained in the encoding's alphabet,
// must not be negative, and must be a rune equal or below '\xff'.
// Padding characters above '\x7f' are encoded as their exact byte value
// rather than using the UTF-8 representation of the codepoint.
func (enc Encoding) WithPadding(padding rune) *Encoding

// Encode encodes src using the encoding enc,
// writing [Encoding.EncodedLen](len(src)) bytes to dst.
//
// The encoding pads the output to a multiple of 8 bytes,
// so Encode is not appropriate for use on individual blocks
// of a large data stream. Use [NewEncoder] instead.
func (enc *Encoding) Encode(dst, src []byte)

// AppendEncode appends the base32 encoded src to dst
// and returns the extended buffer.
func (enc *Encoding) AppendEncode(dst, src []byte) []byte

// EncodeToString returns the base32 encoding of src.
func (enc *Encoding) EncodeToString(src []byte) string

// NewEncoder returns a new base32 stream encoder. Data written to
// the returned writer will be encoded using enc and then written to w.
// Base32 encodings operate in 5-byte blocks; when finished
// writing, the caller must Close the returned encoder to flush any
// partially written blocks.
func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser

// EncodedLen returns the length in bytes of the base32 encoding
// of an input buffer of length n.
func (enc *Encoding) EncodedLen(n int) int

type CorruptInputError int64

func (e CorruptInputError) Error() string

// Decode decodes src using the encoding enc. It writes at most
// [Encoding.DecodedLen](len(src)) bytes to dst and returns the number of bytes
// written. If src contains invalid base32 data, it will return the
// number of bytes successfully written and [CorruptInputError].
// Newline characters (\r and \n) are ignored.
func (enc *Encoding) Decode(dst, src []byte) (n int, err error)

// AppendDecode appends the base32 decoded src to dst
// and returns the extended buffer.
// If the input is malformed, it returns the partially decoded src and an error.
func (enc *Encoding) AppendDecode(dst, src []byte) ([]byte, error)

// DecodeString returns the bytes represented by the base32 string s.
func (enc *Encoding) DecodeString(s string) ([]byte, error)

// NewDecoder constructs a new base32 stream decoder.
func NewDecoder(enc *Encoding, r io.Reader) io.Reader

// DecodedLen returns the maximum length in bytes of the decoded data
// corresponding to n bytes of base32-encoded data.
func (enc *Encoding) DecodedLen(n int) int

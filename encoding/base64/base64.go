// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package base64 implements base64 encoding as specified by RFC 4648.
package base64

import (
	"github.com/shogo82148/std/io"
)

// An Encoding is a radix 64 encoding/decoding scheme, defined by a
// 64-character alphabet. The most common encoding is the "base64"
// encoding defined in RFC 4648 and used in MIME (RFC 2045) and PEM
// (RFC 1421).  RFC 4648 also defines an alternate encoding, which is
// the standard encoding with - and _ substituted for + and /.
type Encoding struct {
	encode    [64]byte
	decodeMap [256]byte
	padChar   rune
	strict    bool
}

const (
	StdPadding rune = '='
	NoPadding  rune = -1
)

// NewEncoding returns a new padded Encoding defined by the given alphabet,
// which must be a 64-byte string that does not contain the padding character
// or CR / LF ('\r', '\n').
// The resulting Encoding uses the default padding character ('='),
// which may be changed or disabled via WithPadding.
func NewEncoding(encoder string) *Encoding

// WithPadding creates a new encoding identical to enc except
// with a specified padding character, or NoPadding to disable padding.
// The padding character must not be '\r' or '\n', must not
// be contained in the encoding's alphabet and must be a rune equal or
// below '\xff'.
func (enc Encoding) WithPadding(padding rune) *Encoding

// Strict creates a new encoding identical to enc except with
// strict decoding enabled. In this mode, the decoder requires that
// trailing padding bits are zero, as described in RFC 4648 section 3.5.
func (enc Encoding) Strict() *Encoding

// StdEncoding is the standard base64 encoding, as defined in
// RFC 4648.
var StdEncoding = NewEncoding(encodeStd)

// URLEncoding is the alternate base64 encoding defined in RFC 4648.
// It is typically used in URLs and file names.
var URLEncoding = NewEncoding(encodeURL)

// RawStdEncoding is the standard raw, unpadded base64 encoding,
// as defined in RFC 4648 section 3.2.
// This is the same as StdEncoding but omits padding characters.
var RawStdEncoding = StdEncoding.WithPadding(NoPadding)

// RawURLEncoding is the unpadded alternate base64 encoding defined in RFC 4648.
// It is typically used in URLs and file names.
// This is the same as URLEncoding but omits padding characters.
var RawURLEncoding = URLEncoding.WithPadding(NoPadding)

// Encode encodes src using the encoding enc, writing
// EncodedLen(len(src)) bytes to dst.
//
// The encoding pads the output to a multiple of 4 bytes,
// so Encode is not appropriate for use on individual blocks
// of a large data stream. Use NewEncoder() instead.
func (enc *Encoding) Encode(dst, src []byte)

// EncodeToString returns the base64 encoding of src.
func (enc *Encoding) EncodeToString(src []byte) string

// NewEncoder returns a new base64 stream encoder. Data written to
// the returned writer will be encoded using enc and then written to w.
// Base64 encodings operate in 4-byte blocks; when finished
// writing, the caller must Close the returned encoder to flush any
// partially written blocks.
func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser

// EncodedLen returns the length in bytes of the base64 encoding
// of an input buffer of length n.
func (enc *Encoding) EncodedLen(n int) int

type CorruptInputError int64

func (e CorruptInputError) Error() string

// Decode decodes src using the encoding enc. It writes at most
// DecodedLen(len(src)) bytes to dst and returns the number of bytes
// written. If src contains invalid base64 data, it will return the
// number of bytes successfully written and CorruptInputError.
// New line characters (\r and \n) are ignored.
func (enc *Encoding) Decode(dst, src []byte) (n int, err error)

// DecodeString returns the bytes represented by the base64 string s.
func (enc *Encoding) DecodeString(s string) ([]byte, error)

// NewDecoder constructs a new base64 stream decoder.
func NewDecoder(enc *Encoding, r io.Reader) io.Reader

// DecodedLen returns the maximum length in bytes of the decoded data
// corresponding to n bytes of base64-encoded data.
func (enc *Encoding) DecodedLen(n int) int

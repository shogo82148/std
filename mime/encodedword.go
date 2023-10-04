// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mime

import (
	"github.com/shogo82148/std/io"
)

// A WordEncoder is an RFC 2047 encoded-word encoder.
type WordEncoder byte

const (
	// BEncoding represents Base64 encoding scheme as defined by RFC 2045.
	BEncoding = WordEncoder('b')
	// QEncoding represents the Q-encoding scheme as defined by RFC 2047.
	QEncoding = WordEncoder('q')
)

// Encode returns the encoded-word form of s. If s is ASCII without special
// characters, it is returned unchanged. The provided charset is the IANA
// charset name of s. It is case insensitive.
func (e WordEncoder) Encode(charset, s string) string

// A WordDecoder decodes MIME headers containing RFC 2047 encoded-words.
type WordDecoder struct {
	// CharsetReader, if non-nil, defines a function to generate
	// charset-conversion readers, converting from the provided
	// charset into UTF-8.
	// Charsets are always lower-case. utf-8, iso-8859-1 and us-ascii charsets
	// are handled by default.
	// One of the CharsetReader's result values must be non-nil.
	CharsetReader func(charset string, input io.Reader) (io.Reader, error)
}

// Decode decodes an RFC 2047 encoded-word.
func (d *WordDecoder) Decode(word string) (string, error)

// DecodeHeader decodes all encoded-words of the given string. It returns an
// error if and only if CharsetReader of d returns an error.
func (d *WordDecoder) DecodeHeader(header string) (string, error)

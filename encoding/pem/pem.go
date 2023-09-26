// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pem implements the PEM data encoding, which originated in Privacy
// Enhanced Mail. The most common use of PEM encoding today is in TLS keys and
// certificates. See RFC 1421.
package pem

import (
	"github.com/shogo82148/std/io"
)

// A Block represents a PEM encoded structure.
//
// The encoded form is:
//
//	-----BEGIN Type-----
//	Headers
//	base64-encoded Bytes
//	-----END Type-----
//
// where Headers is a possibly empty sequence of Key: Value lines.
type Block struct {
	Type    string
	Headers map[string]string
	Bytes   []byte
}

// Decode will find the next PEM formatted block (certificate, private key
// etc) in the input. It returns that block and the remainder of the input. If
// no PEM data is found, p is nil and the whole of the input is returned in
// rest.
func Decode(data []byte) (p *Block, rest []byte)

func Encode(out io.Writer, b *Block) (err error)

func EncodeToMemory(b *Block) []byte

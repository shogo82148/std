// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package png implements a PNG image decoder and encoder.
//
// The PNG specification is at http://www.w3.org/TR/PNG/.
package png

import (
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/io"
)

// Color type, as per the PNG spec.

// A cb is a combination of color type and bit depth.

// Filter type, as per the PNG spec.

// Decoding stage.
// The PNG specification says that the IHDR, PLTE (if present), IDAT and IEND
// chunks must appear in that order. There may be multiple IDAT chunks, and
// IDAT chunks must be sequential (i.e. they may not have any other chunks
// between them).
// http://www.w3.org/TR/PNG/#5ChunkOrdering

// A FormatError reports that the input is not a valid PNG.
type FormatError string

func (e FormatError) Error() string

// An UnsupportedError reports that the input uses a valid but unimplemented PNG feature.
type UnsupportedError string

func (e UnsupportedError) Error() string

// Decode reads a PNG image from r and returns it as an image.Image.
// The type of Image returned depends on the PNG contents.
func Decode(r io.Reader) (image.Image, error)

// DecodeConfig returns the color model and dimensions of a PNG image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (image.Config, error)

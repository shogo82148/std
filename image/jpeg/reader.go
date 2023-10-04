// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package jpeg implements a JPEG image decoder and encoder.
//
// JPEG is defined in ITU-T T.81: https://www.w3.org/Graphics/JPEG/itu-t81.pdf.
package jpeg

import (
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/io"
)

// A FormatError reports that the input is not a valid JPEG.
type FormatError string

func (e FormatError) Error() string

// An UnsupportedError reports that the input uses a valid but unimplemented JPEG feature.
type UnsupportedError string

func (e UnsupportedError) Error() string

// Deprecated: Reader is not used by the image/jpeg package and should
// not be used by others. It is kept for compatibility.
type Reader interface {
	io.ByteReader
	io.Reader
}

// Decode reads a JPEG image from r and returns it as an image.Image.
func Decode(r io.Reader) (image.Image, error)

// DecodeConfig returns the color model and dimensions of a JPEG image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (image.Config, error)

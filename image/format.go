// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// ErrFormat indicates that decoding encountered an unknown format.
var ErrFormat = errors.New("image: unknown format")

// RegisterFormat registers an image format for use by [Decode].
// Name is the name of the format, like "jpeg" or "png".
// Magic is the magic prefix that identifies the format's encoding. The magic
// string can contain "?" wildcards that each match any one byte.
// [Decode] is the function that decodes the encoded image.
// [DecodeConfig] is the function that decodes just its configuration.
func RegisterFormat(name, magic string, decode func(io.Reader) (Image, error), decodeConfig func(io.Reader) (Config, error))

// Decode decodes an image that has been encoded in a registered format.
// The string returned is the format name used during format registration.
// Format registration is typically done by an init function in the codec-
// specific package.
//
// Decoding may allocate memory proportional to the width and height in the
// image header before all pixel data is consumed or validated. When
// decoding untrusted input, call [DecodeConfig] first to inspect dimensions
// and reject images that would exceed resource limits; see the "Security
// Considerations" section in the [image] package documentation.
func Decode(r io.Reader) (Image, string, error)

// DecodeConfig decodes the color model and dimensions of an image that has
// been encoded in a registered format. The string returned is the format name
// used during format registration. Format registration is typically done by
// an init function in the codec-specific package.
//
// DecodeConfig reads only format headers and does not allocate a full-size
// pixel buffer, so it can be used to check dimensions before calling [Decode].
func DecodeConfig(r io.Reader) (Config, string, error)

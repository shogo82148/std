// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gif

import (
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/image/draw"
	"github.com/shogo82148/std/io"
)

// Graphic control extension fields.

// writer is a buffered writer.

// encoder encodes an image to the GIF format.

// blockWriter writes the block structure of GIF image data, which
// comprises (n, (n bytes)) blocks, with 1 <= n <= 255. It is the
// writer given to the LZW encoder, which is thus immune to the
// blocking.

// Options are the encoding parameters.
type Options struct {
	NumColors int

	Quantizer draw.Quantizer

	Drawer draw.Drawer
}

// EncodeAll writes the images in g to w in GIF format with the
// given loop count and delay between frames.
func EncodeAll(w io.Writer, g *GIF) error

// Encode writes the Image m to w in GIF format.
func Encode(w io.Writer, m image.Image, o *Options) error

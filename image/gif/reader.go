// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gif implements a GIF image decoder and encoder.
//
// The GIF specification is at https://www.w3.org/Graphics/GIF/spec-gif89a.txt.
package gif

import (
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/io"
)

// If the io.Reader does not also have ReadByte, then decode will introduce its own buffering.

// Masks etc.

// Disposal Methods.
const (
	DisposalNone       = 0x01
	DisposalBackground = 0x02
	DisposalPrevious   = 0x03
)

// Section indicators.

// Extensions.

// decoder is the type used to decode a GIF file.

// blockReader parses the block structure of GIF image data, which comprises
// (n, (n bytes)) blocks, with 1 <= n <= 255. It is the reader given to the
// LZW decoder, which is thus immune to the blocking. After the LZW decoder
// completes, there will be a 0-byte block remaining (0, ()), which is
// consumed when checking that the blockReader is exhausted.
//
// To avoid the allocation of a bufio.Reader for the lzw Reader, blockReader
// implements io.ByteReader and buffers blocks into the decoder's "tmp" buffer.

// interlaceScan defines the ordering for a pass of the interlace algorithm.

// interlacing represents the set of scans in an interlaced GIF image.

// Decode reads a GIF image from r and returns the first embedded
// image as an image.Image.
func Decode(r io.Reader) (image.Image, error)

// GIF represents the possibly multiple images stored in a GIF file.
type GIF struct {
	Image []*image.Paletted
	Delay []int

	LoopCount int

	Disposal []byte

	Config image.Config

	BackgroundIndex byte
}

// DecodeAll reads a GIF image from r and returns the sequential frames
// and timing information.
func DecodeAll(r io.Reader) (*GIF, error)

// DecodeConfig returns the global color model and dimensions of a GIF image
// without decoding the entire image.
func DecodeConfig(r io.Reader) (image.Config, error)

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jpeg

import (
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/io"
)

// bitCount counts the number of bits needed to hold an integer.

// unscaledQuant are the unscaled quantization tables in zig-zag order. Each
// encoder copies and scales the tables according to its quality parameter.
// The values are derived from section K.1 after converting from natural to
// zig-zag order.

// huffmanSpec specifies a Huffman encoding.

// theHuffmanSpec is the Huffman encoding specifications.
// This encoder uses the same Huffman encoding for all images.

// huffmanLUT is a compiled look-up table representation of a huffmanSpec.
// Each value maps to a uint32 of which the 8 most significant bits hold the
// codeword size in bits and the 24 least significant bits hold the codeword.
// The maximum codeword size is 16 bits.

// theHuffmanLUT are compiled representations of theHuffmanSpec.

// writer is a buffered writer.

// encoder encodes an image to the JPEG format.

// sosHeaderY is the SOS marker "\xff\xda" followed by 8 bytes:
//	- the marker length "\x00\x08",
//	- the number of components "\x01",
//	- component 1 uses DC table 0 and AC table 0 "\x01\x00",
//	- the bytes "\x00\x3f\x00". Section B.2.3 of the spec says that for
//	  sequential DCTs, those bytes (8-bit Ss, 8-bit Se, 4-bit Ah, 4-bit Al)
//	  should be 0x00, 0x3f, 0x00<<4 | 0x00.

// sosHeaderYCbCr is the SOS marker "\xff\xda" followed by 12 bytes:
//	- the marker length "\x00\x0c",
//	- the number of components "\x03",
//	- component 1 uses DC table 0 and AC table 0 "\x01\x00",
//	- component 2 uses DC table 1 and AC table 1 "\x02\x11",
//	- component 3 uses DC table 1 and AC table 1 "\x03\x11",
//	- the bytes "\x00\x3f\x00". Section B.2.3 of the spec says that for
//	  sequential DCTs, those bytes (8-bit Ss, 8-bit Se, 4-bit Ah, 4-bit Al)
//	  should be 0x00, 0x3f, 0x00<<4 | 0x00.

// DefaultQuality is the default quality encoding parameter.
const DefaultQuality = 75

// Options are the encoding parameters.
// Quality ranges from 1 to 100 inclusive, higher is better.
type Options struct {
	Quality int
}

// Encode writes the Image m to w in JPEG 4:2:0 baseline format with the given
// options. Default parameters are used if a nil *Options is passed.
func Encode(w io.Writer, m image.Image, o *Options) error

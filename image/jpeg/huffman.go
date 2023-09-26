// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jpeg

// Each code is at most 16 bits long.

// Each decoded value is a uint8, so there are at most 256 such values.

// Bit stream for the Huffman decoder.
// The n least significant bits of a form the unread bits, to be read in MSB to LSB order.

// Huffman table decoder, specified in section C.

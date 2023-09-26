// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jpeg

// maxCodeLength is the maximum (inclusive) number of bits in a Huffman code.

// maxNCodes is the maximum (inclusive) number of codes in a Huffman tree.

// lutSize is the log-2 size of the Huffman decoder's look-up table.

// huffman is a Huffman decoder, specified in section C.

// errShortHuffmanData means that an unexpected EOF occurred while decoding
// Huffman data.

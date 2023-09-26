// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bzip2

// A huffmanTree is a binary tree which is navigated, bit-by-bit to reach a
// symbol.

// A huffmanNode is a node in the tree. left and right contain indexes into the
// nodes slice of the tree. If left or right is invalidNodeValue then the child
// is a left node and its value is in leftValue/rightValue.
//
// The symbols are uint16s because bzip2 encodes not only MTF indexes in the
// tree, but also two magic values for run-length encoding and an EOF symbol.
// Thus there are more than 256 possible symbols.

// invalidNodeValue is an invalid index which marks a leaf node in the tree.

// huffmanSymbolLengthPair contains a symbol and its code length.

// huffmanSymbolLengthPair is used to provide an interface for sorting.

// huffmanCode contains a symbol, its code and code length.

// huffmanCodes is used to provide an interface for sorting.

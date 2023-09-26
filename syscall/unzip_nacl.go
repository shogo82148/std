// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Small in-memory unzip implementation.
// A simplified copy of the pre-Go 1 compress/flate/inflate.go
// and a modified copy of the zip reader in package time.
// (The one in package time does not support decompression; this one does.)

package syscall

// Hard-coded Huffman tables for DEFLATE algorithm.
// See RFC 1951, section 3.2.6.

// Huffman decoder is based on
// J. Brian Connell, ``A Huffman-Shannon-Fano Code,''
// Proceedings of the IEEE, 61(7) (July 1973), pp 1046-1047.

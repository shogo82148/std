// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flate

// dictDecoder implements the LZ77 sliding dictionary as used in decompression.
// LZ77 decompresses data through sequences of two forms of commands:
//
//	* Literal insertions: Runs of one or more symbols are inserted into the data
//	stream as is. This is accomplished through the writeByte method for a
//	single symbol, or combinations of writeSlice/writeMark for multiple symbols.
//	Any valid stream must start with a literal insertion if no preset dictionary
//	is used.
//
//	* Backward copies: Runs of one or more symbols are copied from previously
//	emitted data. Backward copies come as the tuple (dist, length) where dist
//	determines how far back in the stream to copy from and length determines how
//	many bytes to copy. Note that it is valid for the length to be greater than
//	the distance. Since LZ77 uses forward copies, that situation is used to
//	perform a form of run-length encoding on repeated runs of symbols.
//	The writeCopy and tryWriteCopy are used to implement this command.
//
// For performance reasons, this implementation performs little to no sanity
// checks about the arguments. As such, the invariants documented for each
// method call must be respected.

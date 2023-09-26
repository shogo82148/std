// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob

// decoderState is the execution state of an instance of the decoder. A new state
// is created for nested objects.

// decOp is the signature of a decoding operator for a given type.

// The 'instructions' of the decoding machine

// The encoder engine is an array of instructions indexed by field number of the incoming
// decoder.  It is executed with random access according to field number.

// Index by Go types.

// Indexed by gob types.  tComplex will be added during type.init().

// emptyStruct is the type we compile into when ignoring a struct value.

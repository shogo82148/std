// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// inlinedCall is the encoding of entries in the FUNCDATA_InlTree table.

// An inlineUnwinder iterates over the stack of inlined calls at a PC by
// decoding the inline table. The last step of iteration is always the frame of
// the physical function, so there's always at least one frame.
//
// This is typically used as:
//
//	for u, uf := newInlineUnwinder(...); uf.valid(); uf = u.next(uf) { ... }
//
// Implementation note: This is used in contexts that disallow write barriers.
// Hence, the constructor returns this by value and pointer receiver methods
// must not mutate pointer fields. Also, we keep the mutable state in a separate
// struct mostly to keep both structs SSA-able, which generates much better
// code.

// An inlineFrame is a position in an inlineUnwinder.

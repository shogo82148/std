// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasm.stackswitch

// Continuation table for WebAssembly stack switching (GOWASM=stackswitch).
//
// With stack switching, a g's WebAssembly execution stack is a continuation,
// a reference of type (ref null $cont). Reference values are opaque and
// cannot be stored in linear memory, so they cannot live in a g. Instead the
// linker declares a continuation table, wasm table wasmContTable, and every
// g owns one slot of it, recorded in g.wasmCont.index. A slot holds the g's
// suspended continuation, or null if the g has not started yet.
//
// m0.g0 is allocated statically rather than by malg, so its index is the zero
// value: slot wasmContG0Slot belongs to it. malg gives every other g a slot
// of its own, starting at 1, which the g and its continuation keep for their
// existence, including while the g is on a free list. Like the g's
// themselves, slots are never freed, so the table grows to the number of g's
// ever allocated.
//
// wasm runs a single thread, so the allocation state below needs no locking.

package runtime

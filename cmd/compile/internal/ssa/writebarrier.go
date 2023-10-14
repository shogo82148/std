// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// A ZeroRegion records parts of an object which are known to be zero.
// A ZeroRegion only applies to a single memory state.
// Each bit in mask is set if the corresponding pointer-sized word of
// the base object is known to be zero.
// In other words, if mask & (1<<i) != 0, then [base+i*ptrSize, base+(i+1)*ptrSize)
// is known to be zero.
type ZeroRegion struct {
	base *Value
	mask uint64
}

// IsStackAddr reports whether v is known to be an address of a stack slot.
func IsStackAddr(v *Value) bool

// IsGlobalAddr reports whether v is known to be an address of a global (or nil).
func IsGlobalAddr(v *Value) bool

// IsReadOnlyGlobalAddr reports whether v is known to be an address of a read-only global.
func IsReadOnlyGlobalAddr(v *Value) bool

// IsNewObject reports whether v is a pointer to a freshly allocated & zeroed object,
// if so, also returns the memory state mem at which v is zero.
func IsNewObject(v *Value, select1 []*Value) (mem *Value, ok bool)

// IsSanitizerSafeAddr reports whether v is known to be an address
// that doesn't need instrumentation.
func IsSanitizerSafeAddr(v *Value) bool

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

import "github.com/shogo82148/std/unsafe"

// Int32 is an atomically accessed int32 value.
//
// An Int32 must not be copied.
type Int32 struct {
	noCopy noCopy
	value  int32
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (i *Int32) Load() int32

// Store updates the value atomically.
//
//go:nosplit
func (i *Int32) Store(value int32)

// CompareAndSwap atomically compares i's value with old,
// and if they're equal, swaps i's value with new.
// It reports whether the swap ran.
//
//go:nosplit
func (i *Int32) CompareAndSwap(old, new int32) bool

// Swap replaces i's value with new, returning
// i's value before the replacement.
//
//go:nosplit
func (i *Int32) Swap(new int32) int32

// Add adds delta to i atomically, returning
// the new updated value.
//
// This operation wraps around in the usual
// two's-complement way.
//
//go:nosplit
func (i *Int32) Add(delta int32) int32

// Int64 is an atomically accessed int64 value.
//
// 8-byte aligned on all platforms, unlike a regular int64.
//
// An Int64 must not be copied.
type Int64 struct {
	noCopy noCopy
	_      align64
	value  int64
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (i *Int64) Load() int64

// Store updates the value atomically.
//
//go:nosplit
func (i *Int64) Store(value int64)

// CompareAndSwap atomically compares i's value with old,
// and if they're equal, swaps i's value with new.
// It reports whether the swap ran.
//
//go:nosplit
func (i *Int64) CompareAndSwap(old, new int64) bool

// Swap replaces i's value with new, returning
// i's value before the replacement.
//
//go:nosplit
func (i *Int64) Swap(new int64) int64

// Add adds delta to i atomically, returning
// the new updated value.
//
// This operation wraps around in the usual
// two's-complement way.
//
//go:nosplit
func (i *Int64) Add(delta int64) int64

// Uint8 is an atomically accessed uint8 value.
//
// A Uint8 must not be copied.
type Uint8 struct {
	noCopy noCopy
	value  uint8
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (u *Uint8) Load() uint8

// Store updates the value atomically.
//
//go:nosplit
func (u *Uint8) Store(value uint8)

// And takes value and performs a bit-wise
// "and" operation with the value of u, storing
// the result into u.
//
// The full process is performed atomically.
//
//go:nosplit
func (u *Uint8) And(value uint8)

// Or takes value and performs a bit-wise
// "or" operation with the value of u, storing
// the result into u.
//
// The full process is performed atomically.
//
//go:nosplit
func (u *Uint8) Or(value uint8)

// Bool is an atomically accessed bool value.
//
// A Bool must not be copied.
type Bool struct {
	// Inherits noCopy from Uint8.
	u Uint8
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (b *Bool) Load() bool

// Store updates the value atomically.
//
//go:nosplit
func (b *Bool) Store(value bool)

// Uint32 is an atomically accessed uint32 value.
//
// A Uint32 must not be copied.
type Uint32 struct {
	noCopy noCopy
	value  uint32
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (u *Uint32) Load() uint32

// LoadAcquire is a partially unsynchronized version
// of Load that relaxes ordering constraints. Other threads
// may observe operations that precede this operation to
// occur after it, but no operation that occurs after it
// on this thread can be observed to occur before it.
//
// WARNING: Use sparingly and with great care.
//
//go:nosplit
func (u *Uint32) LoadAcquire() uint32

// Store updates the value atomically.
//
//go:nosplit
func (u *Uint32) Store(value uint32)

// StoreRelease is a partially unsynchronized version
// of Store that relaxes ordering constraints. Other threads
// may observe operations that occur after this operation to
// precede it, but no operation that precedes it
// on this thread can be observed to occur after it.
//
// WARNING: Use sparingly and with great care.
//
//go:nosplit
func (u *Uint32) StoreRelease(value uint32)

// CompareAndSwap atomically compares u's value with old,
// and if they're equal, swaps u's value with new.
// It reports whether the swap ran.
//
//go:nosplit
func (u *Uint32) CompareAndSwap(old, new uint32) bool

// CompareAndSwapRelease is a partially unsynchronized version
// of Cas that relaxes ordering constraints. Other threads
// may observe operations that occur after this operation to
// precede it, but no operation that precedes it
// on this thread can be observed to occur after it.
// It reports whether the swap ran.
//
// WARNING: Use sparingly and with great care.
//
//go:nosplit
func (u *Uint32) CompareAndSwapRelease(old, new uint32) bool

// Swap replaces u's value with new, returning
// u's value before the replacement.
//
//go:nosplit
func (u *Uint32) Swap(value uint32) uint32

// And takes value and performs a bit-wise
// "and" operation with the value of u, storing
// the result into u.
//
// The full process is performed atomically.
//
//go:nosplit
func (u *Uint32) And(value uint32)

// Or takes value and performs a bit-wise
// "or" operation with the value of u, storing
// the result into u.
//
// The full process is performed atomically.
//
//go:nosplit
func (u *Uint32) Or(value uint32)

// Add adds delta to u atomically, returning
// the new updated value.
//
// This operation wraps around in the usual
// two's-complement way.
//
//go:nosplit
func (u *Uint32) Add(delta int32) uint32

// Uint64 is an atomically accessed uint64 value.
//
// 8-byte aligned on all platforms, unlike a regular uint64.
//
// A Uint64 must not be copied.
type Uint64 struct {
	noCopy noCopy
	_      align64
	value  uint64
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (u *Uint64) Load() uint64

// Store updates the value atomically.
//
//go:nosplit
func (u *Uint64) Store(value uint64)

// CompareAndSwap atomically compares u's value with old,
// and if they're equal, swaps u's value with new.
// It reports whether the swap ran.
//
//go:nosplit
func (u *Uint64) CompareAndSwap(old, new uint64) bool

// Swap replaces u's value with new, returning
// u's value before the replacement.
//
//go:nosplit
func (u *Uint64) Swap(value uint64) uint64

// Add adds delta to u atomically, returning
// the new updated value.
//
// This operation wraps around in the usual
// two's-complement way.
//
//go:nosplit
func (u *Uint64) Add(delta int64) uint64

// Uintptr is an atomically accessed uintptr value.
//
// A Uintptr must not be copied.
type Uintptr struct {
	noCopy noCopy
	value  uintptr
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (u *Uintptr) Load() uintptr

// LoadAcquire is a partially unsynchronized version
// of Load that relaxes ordering constraints. Other threads
// may observe operations that precede this operation to
// occur after it, but no operation that occurs after it
// on this thread can be observed to occur before it.
//
// WARNING: Use sparingly and with great care.
//
//go:nosplit
func (u *Uintptr) LoadAcquire() uintptr

// Store updates the value atomically.
//
//go:nosplit
func (u *Uintptr) Store(value uintptr)

// StoreRelease is a partially unsynchronized version
// of Store that relaxes ordering constraints. Other threads
// may observe operations that occur after this operation to
// precede it, but no operation that precedes it
// on this thread can be observed to occur after it.
//
// WARNING: Use sparingly and with great care.
//
//go:nosplit
func (u *Uintptr) StoreRelease(value uintptr)

// CompareAndSwap atomically compares u's value with old,
// and if they're equal, swaps u's value with new.
// It reports whether the swap ran.
//
//go:nosplit
func (u *Uintptr) CompareAndSwap(old, new uintptr) bool

// Swap replaces u's value with new, returning
// u's value before the replacement.
//
//go:nosplit
func (u *Uintptr) Swap(value uintptr) uintptr

// Add adds delta to u atomically, returning
// the new updated value.
//
// This operation wraps around in the usual
// two's-complement way.
//
//go:nosplit
func (u *Uintptr) Add(delta uintptr) uintptr

// Float64 is an atomically accessed float64 value.
//
// 8-byte aligned on all platforms, unlike a regular float64.
//
// A Float64 must not be copied.
type Float64 struct {
	// Inherits noCopy and align64 from Uint64.
	u Uint64
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (f *Float64) Load() float64

// Store updates the value atomically.
//
//go:nosplit
func (f *Float64) Store(value float64)

// UnsafePointer is an atomically accessed unsafe.Pointer value.
//
// Note that because of the atomicity guarantees, stores to values
// of this type never trigger a write barrier, and the relevant
// methods are suffixed with "NoWB" to indicate that explicitly.
// As a result, this type should be used carefully, and sparingly,
// mostly with values that do not live in the Go heap anyway.
//
// An UnsafePointer must not be copied.
type UnsafePointer struct {
	noCopy noCopy
	value  unsafe.Pointer
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (u *UnsafePointer) Load() unsafe.Pointer

// StoreNoWB updates the value atomically.
//
// WARNING: As the name implies this operation does *not*
// perform a write barrier on value, and so this operation may
// hide pointers from the GC. Use with care and sparingly.
// It is safe to use with values not found in the Go heap.
// Prefer Store instead.
//
//go:nosplit
func (u *UnsafePointer) StoreNoWB(value unsafe.Pointer)

// Store updates the value atomically.
func (u *UnsafePointer) Store(value unsafe.Pointer)

// CompareAndSwapNoWB atomically (with respect to other methods)
// compares u's value with old, and if they're equal,
// swaps u's value with new.
// It reports whether the swap ran.
//
// WARNING: As the name implies this operation does *not*
// perform a write barrier on value, and so this operation may
// hide pointers from the GC. Use with care and sparingly.
// It is safe to use with values not found in the Go heap.
// Prefer CompareAndSwap instead.
//
//go:nosplit
func (u *UnsafePointer) CompareAndSwapNoWB(old, new unsafe.Pointer) bool

// CompareAndSwap atomically compares u's value with old,
// and if they're equal, swaps u's value with new.
// It reports whether the swap ran.
func (u *UnsafePointer) CompareAndSwap(old, new unsafe.Pointer) bool

// Pointer is an atomic pointer of type *T.
type Pointer[T any] struct {
	u UnsafePointer
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (p *Pointer[T]) Load() *T

// StoreNoWB updates the value atomically.
//
// WARNING: As the name implies this operation does *not*
// perform a write barrier on value, and so this operation may
// hide pointers from the GC. Use with care and sparingly.
// It is safe to use with values not found in the Go heap.
// Prefer Store instead.
//
//go:nosplit
func (p *Pointer[T]) StoreNoWB(value *T)

// Store updates the value atomically.
//
//go:nosplit
func (p *Pointer[T]) Store(value *T)

// CompareAndSwapNoWB atomically (with respect to other methods)
// compares u's value with old, and if they're equal,
// swaps u's value with new.
// It reports whether the swap ran.
//
// WARNING: As the name implies this operation does *not*
// perform a write barrier on value, and so this operation may
// hide pointers from the GC. Use with care and sparingly.
// It is safe to use with values not found in the Go heap.
// Prefer CompareAndSwap instead.
//
//go:nosplit
func (p *Pointer[T]) CompareAndSwapNoWB(old, new *T) bool

// CompareAndSwap atomically (with respect to other methods)
// compares u's value with old, and if they're equal,
// swaps u's value with new.
// It reports whether the swap ran.
func (p *Pointer[T]) CompareAndSwap(old, new *T) bool

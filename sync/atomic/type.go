// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

import "github.com/shogo82148/std/unsafe"

// A Bool is an atomic boolean value.
// The zero value is false.
type Bool struct {
	_ noCopy
	v uint32
}

// Load atomically loads and returns the value stored in x.
func (x *Bool) Load() bool

// Store atomically stores val into x.
func (x *Bool) Store(val bool)

// Swap atomically stores new into x and returns the previous value.
func (x *Bool) Swap(new bool) (old bool)

// CompareAndSwap executes the compare-and-swap operation for the boolean value x.
func (x *Bool) CompareAndSwap(old, new bool) (swapped bool)

// For testing *Pointer[T]'s methods can be inlined.
// Keep in sync with cmd/compile/internal/test/inl_test.go:TestIntendedInlining.
var _ = &Pointer[int]{}

// A Pointer is an atomic pointer of type *T. The zero value is a nil *T.
type Pointer[T any] struct {
	_ [0]*T

	_ noCopy
	v unsafe.Pointer
}

// Load atomically loads and returns the value stored in x.
func (x *Pointer[T]) Load() *T

// Store atomically stores val into x.
func (x *Pointer[T]) Store(val *T)

// Swap atomically stores new into x and returns the previous value.
func (x *Pointer[T]) Swap(new *T) (old *T)

// CompareAndSwap executes the compare-and-swap operation for x.
func (x *Pointer[T]) CompareAndSwap(old, new *T) (swapped bool)

// An Int32 is an atomic int32. The zero value is zero.
type Int32 struct {
	_ noCopy
	v int32
}

// Load atomically loads and returns the value stored in x.
func (x *Int32) Load() int32

// Store atomically stores val into x.
func (x *Int32) Store(val int32)

// Swap atomically stores new into x and returns the previous value.
func (x *Int32) Swap(new int32) (old int32)

// CompareAndSwap executes the compare-and-swap operation for x.
func (x *Int32) CompareAndSwap(old, new int32) (swapped bool)

// Add atomically adds delta to x and returns the new value.
func (x *Int32) Add(delta int32) (new int32)

// An Int64 is an atomic int64. The zero value is zero.
type Int64 struct {
	_ noCopy
	_ align64
	v int64
}

// Load atomically loads and returns the value stored in x.
func (x *Int64) Load() int64

// Store atomically stores val into x.
func (x *Int64) Store(val int64)

// Swap atomically stores new into x and returns the previous value.
func (x *Int64) Swap(new int64) (old int64)

// CompareAndSwap executes the compare-and-swap operation for x.
func (x *Int64) CompareAndSwap(old, new int64) (swapped bool)

// Add atomically adds delta to x and returns the new value.
func (x *Int64) Add(delta int64) (new int64)

// A Uint32 is an atomic uint32. The zero value is zero.
type Uint32 struct {
	_ noCopy
	v uint32
}

// Load atomically loads and returns the value stored in x.
func (x *Uint32) Load() uint32

// Store atomically stores val into x.
func (x *Uint32) Store(val uint32)

// Swap atomically stores new into x and returns the previous value.
func (x *Uint32) Swap(new uint32) (old uint32)

// CompareAndSwap executes the compare-and-swap operation for x.
func (x *Uint32) CompareAndSwap(old, new uint32) (swapped bool)

// Add atomically adds delta to x and returns the new value.
func (x *Uint32) Add(delta uint32) (new uint32)

// A Uint64 is an atomic uint64. The zero value is zero.
type Uint64 struct {
	_ noCopy
	_ align64
	v uint64
}

// Load atomically loads and returns the value stored in x.
func (x *Uint64) Load() uint64

// Store atomically stores val into x.
func (x *Uint64) Store(val uint64)

// Swap atomically stores new into x and returns the previous value.
func (x *Uint64) Swap(new uint64) (old uint64)

// CompareAndSwap executes the compare-and-swap operation for x.
func (x *Uint64) CompareAndSwap(old, new uint64) (swapped bool)

// Add atomically adds delta to x and returns the new value.
func (x *Uint64) Add(delta uint64) (new uint64)

// A Uintptr is an atomic uintptr. The zero value is zero.
type Uintptr struct {
	_ noCopy
	v uintptr
}

// Load atomically loads and returns the value stored in x.
func (x *Uintptr) Load() uintptr

// Store atomically stores val into x.
func (x *Uintptr) Store(val uintptr)

// Swap atomically stores new into x and returns the previous value.
func (x *Uintptr) Swap(new uintptr) (old uintptr)

// CompareAndSwap executes the compare-and-swap operation for x.
func (x *Uintptr) CompareAndSwap(old, new uintptr) (swapped bool)

// Add atomically adds delta to x and returns the new value.
func (x *Uintptr) Add(delta uintptr) (new uintptr)

// noCopy may be added to structs which must not be copied
// after the first use.
//
// See https://golang.org/issues/8005#issuecomment-190753527
// for details.
//
// Note that it must not be embedded, due to the Lock and Unlock methods.

// align64 may be added to structs that must be 64-bit aligned.
// This struct is recognized by a special case in the compiler
// and will not work if copied to any other package.

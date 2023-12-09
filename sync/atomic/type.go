// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

import "github.com/shogo82148/std/unsafe"

// Boolはアトミックなブーリアン値です。
// ゼロ値はfalseです。
type Bool struct {
	_ noCopy
	v uint32
}

// Loadはアトミックにxに格納されている値をロードして返します。
func (x *Bool) Load() bool

// Storeはアトミックにvalをxに格納します。
func (x *Bool) Store(val bool)

// Swapはアトミックにnewをxに格納し、前の値を返します。
func (x *Bool) Swap(new bool) (old bool)

// CompareAndSwapは、ブール値xの比較交換操作を実行します。
func (x *Bool) CompareAndSwap(old, new bool) (swapped bool)

// For testing *Pointer[T]'s methods can be inlined.
// Keep in sync with cmd/compile/internal/test/inl_test.go:TestIntendedInlining.
var _ = &Pointer[int]{}

// Pointerはタイプ*Tのアトミックポインタです。ゼロ値はnil *Tです。
type Pointer[T any] struct {
	// Mention *T in a field to disallow conversion between Pointer types.
	// See go.dev/issue/56603 for more details.
	// Use *T, not T, to avoid spurious recursive type definition errors.
	_ [0]*T

	_ noCopy
	v unsafe.Pointer
}

// Loadはアトミックにxに格納されている値をロードして返します。
func (x *Pointer[T]) Load() *T

// Storeはアトミックにvalをxに格納します。
func (x *Pointer[T]) Store(val *T)

// Swapはアトミックにnewをxに格納し、前の値を返します。
func (x *Pointer[T]) Swap(new *T) (old *T)

// CompareAndSwapは、ポインタxの比較交換操作を実行します。
func (x *Pointer[T]) CompareAndSwap(old, new *T) (swapped bool)

// Int32はアトミックなint32です。ゼロ値はゼロです。
type Int32 struct {
	_ noCopy
	v int32
}

// Loadはアトミックにxに格納されている値をロードして返します。
func (x *Int32) Load() int32

// Storeはアトミックにvalをxに格納します。
func (x *Int32) Store(val int32)

// Swapはアトミックにnewをxに格納し、前の値を返します。
func (x *Int32) Swap(new int32) (old int32)

// CompareAndSwapは、xの比較交換操作を実行します。
func (x *Int32) CompareAndSwap(old, new int32) (swapped bool)

// Addはアトミックにdeltaをxに加え、新しい値を返します。
func (x *Int32) Add(delta int32) (new int32)

// Int64はアトミックなint64です。ゼロ値はゼロです。
type Int64 struct {
	_ noCopy
	_ align64
	v int64
}

// Loadはアトミックにxに格納されている値をロードして返します。
func (x *Int64) Load() int64

// Storeはアトミックにvalをxに格納します。
func (x *Int64) Store(val int64)

// Swapはアトミックにnewをxに格納し、前の値を返します。
func (x *Int64) Swap(new int64) (old int64)

// CompareAndSwapは、xの比較交換操作を実行します。
func (x *Int64) CompareAndSwap(old, new int64) (swapped bool)

// Addはアトミックにdeltaをxに加え、新しい値を返します。
func (x *Int64) Add(delta int64) (new int64)

// Uint32はアトミックなuint32です。ゼロ値はゼロです。
type Uint32 struct {
	_ noCopy
	v uint32
}

// Loadはアトミックにxに格納されている値をロードして返します。
func (x *Uint32) Load() uint32

// Storeはアトミックにvalをxに格納します。
func (x *Uint32) Store(val uint32)

// Swapはアトミックにnewをxに格納し、前の値を返します。
func (x *Uint32) Swap(new uint32) (old uint32)

// CompareAndSwapは、xの比較交換操作を実行します。
func (x *Uint32) CompareAndSwap(old, new uint32) (swapped bool)

// Addはアトミックにdeltaをxに加え、新しい値を返します。
func (x *Uint32) Add(delta uint32) (new uint32)

// Uint64はアトミックなuint64です。ゼロ値はゼロです。
type Uint64 struct {
	_ noCopy
	_ align64
	v uint64
}

// Loadはアトミックにxに格納されている値をロードして返します。
func (x *Uint64) Load() uint64

// Storeはアトミックにvalをxに格納します。
func (x *Uint64) Store(val uint64)

// Swapはアトミックにnewをxに格納し、前の値を返します。
func (x *Uint64) Swap(new uint64) (old uint64)

// CompareAndSwapは、xの比較交換操作を実行します。
func (x *Uint64) CompareAndSwap(old, new uint64) (swapped bool)

// Addはアトミックにdeltaをxに加え、新しい値を返します。
func (x *Uint64) Add(delta uint64) (new uint64)

// Uintptrはアトミックなuintptrです。ゼロ値はゼロです。
type Uintptr struct {
	_ noCopy
	v uintptr
}

// Loadはアトミックにxに格納されている値をロードして返します。
func (x *Uintptr) Load() uintptr

// Storeはアトミックにvalをxに格納します。
func (x *Uintptr) Store(val uintptr)

// Swapはアトミックにnewをxに格納し、前の値を返します。
func (x *Uintptr) Swap(new uintptr) (old uintptr)

// CompareAndSwapは、xの比較交換操作を実行します。
func (x *Uintptr) CompareAndSwap(old, new uintptr) (swapped bool)

// Addはアトミックにdeltaをxに加え、新しい値を返します。
func (x *Uintptr) Add(delta uintptr) (new uintptr)

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージatomicは、同期アルゴリズムの実装に役立つ
// 低レベルのアトミックメモリプリミティブを提供します。
//
// これらの関数は、正しく使用するためには非常に注意が必要です。
// 特別な低レベルのアプリケーションを除き、同期はチャネルや[sync]パッケージの機能を
// 使用して行う方が良いです。
// メモリを共有するために通信を行い、
// メモリを共有するために通信を行わないでください。
//
// SwapT関数によって実装されるスワップ操作は、アトミックの
// 相当するものです：
//
//	old = *addr
//	*addr = new
//	return old
//
// CompareAndSwapT関数によって実装される比較交換操作は、アトミックの
// 相当するものです：
//
//	if *addr == old {
//		*addr = new
//		return true
//	}
//	return false
//
// AddT関数によって実装される加算操作は、アトミックの
// 相当するものです：
//
//	*addr += delta
//	return *addr
//
// LoadTおよびStoreT関数によって実装されるロードおよびストア操作は、
// "return *addr"および"*addr = val"のアトミック相当です。
//
// Goのメモリモデルの用語では、アトミック操作Aの効果が
// アトミック操作Bによって観察される場合、AはBの前に「同期する」。
// さらに、プログラムで実行されるすべてのアトミック操作は、
// あたかも一貫した順序で実行されるかのように振る舞います。
// この定義は、C++の一貫性のあるアトミックとJavaのvolatile変数と
// 同じセマンティクスを提供します。
package atomic

import (
	"github.com/shogo82148/std/unsafe"
)

// SwapInt32はアトミックに新しい値を*addrに格納し、前の*addrの値を返します。
// より使いやすく、エラーが発生しにくい [Int32.Swap] の使用を検討してください。
func SwapInt32(addr *int32, new int32) (old int32)

// SwapInt64はアトミックに新しい値を*addrに格納し、前の*addrの値を返します。
// より使いやすく、エラーが発生しにくい [Int64.Swap] の使用を検討してください
// （特に32ビットプラットフォームを対象とする場合は、バグセクションを参照してください）。
func SwapInt64(addr *int64, new int64) (old int64)

// SwapUint32はアトミックに新しい値を*addrに格納し、前の*addrの値を返します。
// より使いやすく、エラーが発生しにくい [Uint32.Swap] の使用を検討してください。
func SwapUint32(addr *uint32, new uint32) (old uint32)

// SwapUint64はアトミックに新しい値を*addrに格納し、前の*addrの値を返します。
// より使いやすく、エラーが発生しにくい [Uint64.Swap] の使用を検討してください
// （特に32ビットプラットフォームを対象とする場合は、バグセクションを参照してください）。
func SwapUint64(addr *uint64, new uint64) (old uint64)

// SwapUintptrはアトミックに新しい値を*addrに格納し、前の*addrの値を返します。
// より使いやすく、エラーが発生しにくい [Uintptr.Swap] の使用を検討してください。
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)

// SwapPointerはアトミックに新しい値を*addrに格納し、前の*addrの値を返します。
// より使いやすく、エラーが発生しにくい [Pointer.Swap] の使用を検討してください。
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)

// CompareAndSwapInt32は、int32値のための比較交換操作を実行します。
// より使いやすく、エラーが発生しにくい [Int32.CompareAndSwap] の使用を検討してください。
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)

// CompareAndSwapInt64は、int64値のための比較交換操作を実行します。
// より使いやすく、エラーが発生しにくい [Int64.CompareAndSwap] の使用を検討してください
// （特に32ビットプラットフォームを対象とする場合は、バグセクションを参照してください）。
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)

// CompareAndSwapUint32は、uint32値のための比較交換操作を実行します。
// より使いやすく、エラーが発生しにくい [Uint32.CompareAndSwap] の使用を検討してください。
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)

// CompareAndSwapUint64は、uint64値のための比較交換操作を実行します。
// より使いやすく、エラーが発生しにくい [Uint64.CompareAndSwap] の使用を検討してください
// （特に32ビットプラットフォームを対象とする場合は、バグセクションを参照してください）。
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)

// CompareAndSwapUintptrは、uintptr値のための比較交換操作を実行します。
// より使いやすく、エラーが発生しにくい [Uintptr.CompareAndSwap] の使用を検討してください。
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)

// CompareAndSwapPointerは、unsafe.Pointer値のための比較交換操作を実行します。
// より使いやすく、エラーが発生しにくい [Pointer.CompareAndSwap] の使用を検討してください。
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)

// AddInt32はアトミックにdeltaを*addrに加え、新しい値を返します。
// より使いやすく、エラーが発生しにくい [Int32.Add] の使用を検討してください。
func AddInt32(addr *int32, delta int32) (new int32)

// AddUint32はアトミックにdeltaを*addrに加え、新しい値を返します。
// xから符号付き正の定数値cを減算するには、AddUint32(&x, ^uint32(c-1))を行います。
// 特に、xをデクリメントするには、AddUint32(&x, ^uint32(0))を行います。
// より使いやすく、エラーが発生しにくい [Uint32.Add] の使用を検討してください。
func AddUint32(addr *uint32, delta uint32) (new uint32)

// AddInt64はアトミックにdeltaを*addrに加え、新しい値を返します。
// より使いやすく、エラーが発生しにくい [Int64.Add] の使用を検討してください
// （特に32ビットプラットフォームを対象とする場合は、バグセクションを参照してください）。
func AddInt64(addr *int64, delta int64) (new int64)

// AddUint64はアトミックにdeltaを*addrに加え、新しい値を返します。
// xから符号付き正の定数値cを減算するには、AddUint64(&x, ^uint64(c-1))を行います。
// 特に、xをデクリメントするには、AddUint64(&x, ^uint64(0))を行います。
// より使いやすく、エラーが発生しにくい [Uint64.Add] の使用を検討してください
// （特に32ビットプラットフォームを対象とする場合は、バグセクションを参照してください）。
func AddUint64(addr *uint64, delta uint64) (new uint64)

// AddUintptrはアトミックにdeltaを*addrに加え、新しい値を返します。
// より使いやすく、エラーが発生しにくい [Uintptr.Add] の使用を検討してください。
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)

// AndInt32は、提供されたマスクとしてビットマスクを使用して*addr上でビット単位のAND操作をアトミックに実行し、
// 古い値を返します。
// より使いやすく、エラーが発生しにくい [Int32.And] の使用を検討してください。
func AndInt32(addr *int32, mask int32) (old int32)

// AndUint32は、提供されたマスクとしてビットマスクを使用して*addr上でビット単位のAND操作をアトミックに実行し、
// 古い値を返します。
// より使いやすく、エラーが発生しにくい [Uint32.And] の使用を検討してください。
func AndUint32(addr *uint32, mask uint32) (old uint32)

// AndInt64は、提供されたマスクとしてビットマスクを使用して*addr上でビット単位のAND操作をアトミックに実行し、
// 古い値を返します。
// より使いやすく、エラーが発生しにくい [Int64.And] の使用を検討してください。
func AndInt64(addr *int64, mask int64) (old int64)

// AndUint64は、提供されたマスクとしてビットマスクを使用して*addr上でビット単位のAND操作をアトミックに実行し、
// 古い値を返します。
// より使いやすく、エラーが発生しにくい [Uint64.And] の使用を検討してください。
func AndUint64(addr *uint64, mask uint64) (old uint64)

// AndUintptrは、提供されたマスクとしてビットマスクを使用して*addr上でビット単位のAND操作をアトミックに実行し、
// 古い値を返します。
// より使いやすく、エラーが発生しにくい [Uintptr.And] の使用を検討してください。
func AndUintptr(addr *uintptr, mask uintptr) (old uintptr)

// OrInt32 atomically performs a bitwise OR operation on *addr using the bitmask provided as mask
// and returns the old value.
// Consider using the more ergonomic and less error-prone [Int32.Or] instead.
func OrInt32(addr *int32, mask int32) (old int32)

// OrUint32は、提供されたマスクとしてビットマスクを使用して*addr上でビット単位のOR操作をアトミックに実行し、
// 古い値を返します。
// より使いやすく、エラーが発生しにくい [Uint32.Or] の使用を検討してください。
func OrUint32(addr *uint32, mask uint32) (old uint32)

// OrInt64は、提供されたマスクとしてビットマスクを使用して*addr上でビット単位のOR操作をアトミックに実行し、
// 古い値を返します。
// より使いやすく、エラーが発生しにくい [Int64.Or] の使用を検討してください。
func OrInt64(addr *int64, mask int64) (old int64)

// OrUint64は、提供されたマスクとしてビットマスクを使用して*addr上でビット単位のOR操作をアトミックに実行し、
// 古い値を返します。
// より使いやすく、エラーが発生しにくい [Uint64.Or] の使用を検討してください。
func OrUint64(addr *uint64, mask uint64) (old uint64)

// OrUintptrは、提供されたマスクとしてビットマスクを使用して*addr上でビット単位のOR操作をアトミックに実行し、
// 古い値を返します。
// より使いやすく、エラーが発生しにくい [Uintptr.Or] の使用を検討してください。
func OrUintptr(addr *uintptr, mask uintptr) (old uintptr)

// LoadInt32はアトミックに*addrをロードします。
// より使いやすく、エラーが発生しにくい [Int32.Load] の使用を検討してください。
func LoadInt32(addr *int32) (val int32)

// LoadInt64はアトミックに*addrをロードします。
// より使いやすく、エラーが発生しにくい [Int64.Load] の使用を検討してください
// （特に32ビットプラットフォームを対象とする場合は、バグセクションを参照してください）。
func LoadInt64(addr *int64) (val int64)

// LoadUint32はアトミックに*addrをロードします。
// より使いやすく、エラーが発生しにくい [Uint32.Load] の使用を検討してください。
func LoadUint32(addr *uint32) (val uint32)

// LoadUint64はアトミックに*addrをロードします。
// より使いやすく、エラーが発生しにくい [Uint64.Load] の使用を検討してください
// （特に32ビットプラットフォームを対象とする場合は、バグセクションを参照してください）。
func LoadUint64(addr *uint64) (val uint64)

// LoadUintptrはアトミックに*addrをロードします。
// より使いやすく、エラーが発生しにくい [Uintptr.Load] の使用を検討してください。
func LoadUintptr(addr *uintptr) (val uintptr)

// LoadPointerはアトミックに*addrをロードします。
// より使いやすく、エラーが発生しにくい [Pointer.Load] の使用を検討してください。
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)

// StoreInt32はアトミックにvalを*addrに格納します。
// より使いやすく、エラーが発生しにくい [Int32.Store] の使用を検討してください。
func StoreInt32(addr *int32, val int32)

// StoreInt64はアトミックにvalを*addrに格納します。
// より使いやすく、エラーが発生しにくい [Int64.Store] の使用を検討してください
// （特に32ビットプラットフォームを対象とする場合は、バグセクションを参照してください）。
func StoreInt64(addr *int64, val int64)

// StoreUint32はアトミックにvalを*addrに格納します。
// より使いやすく、エラーが発生しにくい [Uint32.Store] の使用を検討してください。
func StoreUint32(addr *uint32, val uint32)

// StoreUint64はアトミックにvalを*addrに格納します。
// より使いやすく、エラーが発生しにくい [Uint64.Store] の使用を検討してください
// （特に32ビットプラットフォームを対象とする場合は、バグセクションを参照してください）。
func StoreUint64(addr *uint64, val uint64)

// StoreUintptrはアトミックにvalを*addrに格納します。
// より使いやすく、エラーが発生しにくい [Uintptr.Store] の使用を検討してください。
func StoreUintptr(addr *uintptr, val uintptr)

// StorePointerはアトミックにvalを*addrに格納します。
// より使いやすく、エラーが発生しにくい [Pointer.Store] の使用を検討してください。
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)

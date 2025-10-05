// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// atomicパッケージは、同期アルゴリズムの実装に役立つ
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
// [the Go memory model] の用語では、アトミック操作Aの効果が
// アトミック操作Bによって観察される場合、AはBの前に「同期する」。
// さらに、プログラムで実行されるすべてのアトミック操作は、
// あたかも一貫した順序で実行されるかのように振る舞います。
// この定義は、C++の一貫性のあるアトミックとJavaのvolatile変数と
// 同じセマンティクスを提供します。
//
// [the Go memory model]: https://go.dev/ref/mem
package atomic

import (
	"github.com/shogo82148/std/unsafe"
)

// SwapInt32は、*addrにnewをアトミックに格納し、以前の*addrの値を返します。
// より使いやすく、エラーが発生しにくい [Int32.Swap] の使用を検討してください。
//
//go:noescape
func SwapInt32(addr *int32, new int32) (old int32)

// SwapUint32は、*addrにnewをアトミックに格納し、以前の*addrの値を返します。
// より使いやすく、エラーが発生しにくい [Uint32.Swap] の使用を検討してください。
//
//go:noescape
func SwapUint32(addr *uint32, new uint32) (old uint32)

// SwapUintptrは、*addrにnewをアトミックに格納し、以前の*addrの値を返します。
// より使いやすく、エラーが発生しにくい [Uintptr.Swap] の使用を検討してください。
//
//go:noescape
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)

// SwapPointerはアトミックに新しい値を*addrに格納し、前の*addrの値を返します。
// より使いやすく、エラーが発生しにくい [Pointer.Swap] の使用を検討してください。
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)

// CompareAndSwapInt32は、int32値のための比較交換操作を実行します。
// より使いやすく、エラーが発生しにくい [Int32.CompareAndSwap] の使用を検討してください。
//
//go:noescape
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)

// CompareAndSwapUint32は、uint32値のための比較交換操作を実行します。
// より使いやすく、エラーが発生しにくい [Uint32.CompareAndSwap] の使用を検討してください。
//
//go:noescape
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)

// CompareAndSwapUintptrは、uintptr値のための比較交換操作を実行します。
// より使いやすく、エラーが発生しにくい [Uintptr.CompareAndSwap] の使用を検討してください。
//
//go:noescape
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)

// CompareAndSwapPointerは、unsafe.Pointer値のための比較交換操作を実行します。
// より使いやすく、エラーが発生しにくい [Pointer.CompareAndSwap] の使用を検討してください。
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)

// AddInt32は、deltaを*addrにアトミックに加算し、新しい値を返します。
// より使いやすく、エラーが発生しにくい [Int32.Add] の使用を検討してください。
//
//go:noescape
func AddInt32(addr *int32, delta int32) (new int32)

// AddUint32は、deltaを*addrにアトミックに加算し、新しい値を返します。
// xから符号付き正の定数値cを減算するには、AddUint32(&x, ^uint32(c-1))を使用します。
// 特に、xをデクリメントするには、AddUint32(&x, ^uint32(0))を使用します。
// より使いやすく、エラーが発生しにくい [Uint32.Add] の使用を検討してください。
//
//go:noescape
func AddUint32(addr *uint32, delta uint32) (new uint32)

// AddUintptrは、deltaを*addrにアトミックに加算し、新しい値を返します。
// より使いやすく、エラーが発生しにくい [Uintptr.Add] の使用を検討してください。
//
//go:noescape
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)

// AndInt32は、指定されたマスクを使って*addr上でビット単位のAND操作をアトミックに実行し、古い値を返します。
// より使いやすく、エラーが発生しにくい [Int32.And] の使用を検討してください。
//
//go:noescape
func AndInt32(addr *int32, mask int32) (old int32)

// AndUint32は、指定されたマスクを使って*addr上でビット単位のAND操作をアトミックに実行し、古い値を返します。
// より使いやすく、エラーが発生しにくい [Uint32.And] の使用を検討してください。
//
//go:noescape
func AndUint32(addr *uint32, mask uint32) (old uint32)

// AndUintptrは、指定されたマスクを使って*addr上でビット単位のAND操作をアトミックに実行し、古い値を返します。
// より使いやすく、エラーが発生しにくい [Uintptr.And] の使用を検討してください。
//
//go:noescape
func AndUintptr(addr *uintptr, mask uintptr) (old uintptr)

// OrInt32は、指定されたマスクを使って*addr上でビット単位のOR操作をアトミックに実行し、古い値を返します。
// より使いやすく、エラーが発生しにくい [Int32.Or] の使用を検討してください。
//
//go:noescape
func OrInt32(addr *int32, mask int32) (old int32)

// OrUint32は、指定されたマスクを使って*addr上でビット単位のOR操作をアトミックに実行し、古い値を返します。
// より使いやすく、エラーが発生しにくい [Uint32.Or] の使用を検討してください。
//
//go:noescape
func OrUint32(addr *uint32, mask uint32) (old uint32)

// OrUintptrは、指定されたマスクを使って*addr上でビット単位のOR操作をアトミックに実行し、古い値を返します。
// より使いやすく、エラーが発生しにくい [Uintptr.Or] の使用を検討してください。
//
//go:noescape
func OrUintptr(addr *uintptr, mask uintptr) (old uintptr)

// LoadInt32はアトミックに*addrをロードします。
// より使いやすく、エラーが発生しにくい [Int32.Load] の使用を検討してください。
//
//go:noescape
func LoadInt32(addr *int32) (val int32)

// LoadUint32はアトミックに*addrをロードします。
// より使いやすく、エラーが発生しにくい [Uint32.Load] の使用を検討してください。
//
//go:noescape
func LoadUint32(addr *uint32) (val uint32)

// LoadUintptrはアトミックに*addrをロードします。
// より使いやすく、エラーが発生しにくい [Uintptr.Load] の使用を検討してください。
//
//go:noescape
func LoadUintptr(addr *uintptr) (val uintptr)

// LoadPointerはアトミックに*addrをロードします。
// より使いやすく、エラーが発生しにくい [Pointer.Load] の使用を検討してください。
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)

// StoreInt32は、valを*addrにアトミックに格納します。
// より使いやすく、エラーが発生しにくい [Int32.Store] の使用を検討してください。
//
//go:noescape
func StoreInt32(addr *int32, val int32)

// StoreUint32は、valを*addrにアトミックに格納します。
// より使いやすく、エラーが発生しにくい [Uint32.Store] の使用を検討してください。
//
//go:noescape
func StoreUint32(addr *uint32, val uint32)

// StoreUintptrは、valを*addrにアトミックに格納します。
// より使いやすく、エラーが発生しにくい [Uintptr.Store] の使用を検討してください。
//
//go:noescape
func StoreUintptr(addr *uintptr, val uintptr)

// StorePointerはアトミックにvalを*addrに格納します。
// より使いやすく、エラーが発生しにくい [Pointer.Store] の使用を検討してください。
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)

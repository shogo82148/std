// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !(386 || arm || mips || mipsle)

package atomic

// SwapInt64 は *addr に new をアトミックに格納し、以前の *addr の値を返します。
// より人間工学的でエラーが起きにくい [Int64.Swap] の使用を検討してください
// （32 ビットプラットフォームをターゲットにする場合は特にご注意ください。bugs セクションを参照）。
//
//go:noescape
func SwapInt64(addr *int64, new int64) (old int64)

// SwapUint64 は *addr に new をアトミックに格納し、以前の *addr の値を返します。
// より人間工学的でエラーが起きにくい [Uint64.Swap] の使用を検討してください
// （32 ビットプラットフォームをターゲットにする場合は特にご注意ください。bugs セクションを参照）。
//
//go:noescape
func SwapUint64(addr *uint64, new uint64) (old uint64)

// CompareAndSwapInt64 は int64 値に対するコンペアアンドスワップ操作を実行します。
// より人間工学的でエラーが起きにくい [Int64.CompareAndSwap] の使用を検討してください
// （32 ビットプラットフォームをターゲットにする場合は特にご注意ください。bugs セクションを参照）。
//
//go:noescape
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)

// CompareAndSwapUint64 は uint64 値に対するコンペアアンドスワップ操作を実行します。
// より人間工学的でエラーが起きにくい [Uint64.CompareAndSwap] の使用を検討してください
// （32 ビットプラットフォームをターゲットにする場合は特にご注意ください。bugs セクションを参照）。
//
//go:noescape
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)

// AddInt64 は *addr に delta をアトミックに加算し、新しい値を返します。
// より人間工学的でエラーが起きにくい [Int64.Add] の使用を検討してください
// （32 ビットプラットフォームをターゲットにする場合は特にご注意ください。bugs セクションを参照）。
//
//go:noescape
func AddInt64(addr *int64, delta int64) (new int64)

// AddUint64 は *addr に delta をアトミックに加算し、新しい値を返します。
// x から符号付き正の定数値 c を減算するには AddUint64(&x, ^uint64(c-1)) を実行します。
// 特に x をデクリメントするには AddUint64(&x, ^uint64(0)) を実行します。
// より人間工学的でエラーが起きにくい [Uint64.Add] の使用を検討してください
// （32 ビットプラットフォームをターゲットにする場合は特にご注意ください。bugs セクションを参照）。
//
//go:noescape
func AddUint64(addr *uint64, delta uint64) (new uint64)

// AndInt64 は mask で指定されたビットマスクを使って *addr にビットAND演算を
// アトミックに実行し、古い値を返します。
// より人間工学的でエラーが起きにくい [Int64.And] の使用を検討してください。
//
//go:noescape
func AndInt64(addr *int64, mask int64) (old int64)

// AndUint64 は mask で指定されたビットマスクを使って *addr にビットAND演算を
// アトミックに実行し、古い値を返します。
// より人間工学的でエラーが起きにくい [Uint64.And] の使用を検討してください。
//
//go:noescape
func AndUint64(addr *uint64, mask uint64) (old uint64)

// OrInt64 は mask で指定されたビットマスクを使って *addr にビットOR演算を
// アトミックに実行し、古い値を返します。
// より人間工学的でエラーが起きにくい [Int64.Or] の使用を検討してください。
//
//go:noescape
func OrInt64(addr *int64, mask int64) (old int64)

// OrUint64 は mask で指定されたビットマスクを使って *addr にビットOR演算を
// アトミックに実行し、古い値を返します。
// より人間工学的でエラーが起きにくい [Uint64.Or] の使用を検討してください。
//
//go:noescape
func OrUint64(addr *uint64, mask uint64) (old uint64)

// LoadInt64 は *addr をアトミックにロードします。
// より人間工学的でエラーが起きにくい [Int64.Load] の使用を検討してください
// （32 ビットプラットフォームをターゲットにする場合は特にご注意ください。bugs セクションを参照）。
//
//go:noescape
func LoadInt64(addr *int64) (val int64)

// LoadUint64 は *addr をアトミックにロードします。
// より人間工学的でエラーが起きにくい [Uint64.Load] の使用を検討してください
// （32 ビットプラットフォームをターゲットにする場合は特にご注意ください。bugs セクションを参照）。
//
//go:noescape
func LoadUint64(addr *uint64) (val uint64)

// StoreInt64 は val を *addr にアトミックに格納します。
// より人間工学的でエラーが起きにくい [Int64.Store] の使用を検討してください
// （32 ビットプラットフォームをターゲットにする場合は特にご注意ください。bugs セクションを参照）。
//
//go:noescape
func StoreInt64(addr *int64, val int64)

// StoreUint64 は val を *addr にアトミックに格納します。
// より人間工学的でエラーが起きにくい [Uint64.Store] の使用を検討してください
// （32 ビットプラットフォームをターゲットにする場合は特にご注意ください。bugs セクションを参照）。
//
//go:noescape
func StoreUint64(addr *uint64, val uint64)

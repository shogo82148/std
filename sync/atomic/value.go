// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

// Valueは、一貫した型の値のアトミックなロードとストアを提供します。
// Valueのゼロ値は [Value.Load] からnilを返します。
// [Value.Store] が呼び出された後、Valueはコピーしてはなりません。
//
// 最初の使用後、Valueはコピーしてはなりません。
type Value struct {
	v any
}

// Loadは、最も最近のStoreによって設定された値を返します。
// このValueに対してStoreの呼び出しがない場合、nilを返します。
func (v *Value) Load() (val any)

// Storeは、[Value] vの値をvalに設定します。
// 与えられたValueに対するStoreのすべての呼び出しは、同じ具体的な型の値を使用しなければなりません。
// 不一致の型をStoreするとパニックを引き起こし、Store(nil)も同様です。
func (v *Value) Store(val any)

// Swapは新しい値をValueに格納し、前の値を返します。Valueが空の場合はnilを返します。
//
// 与えられたValueに対するSwapのすべての呼び出しは、同じ具体的な型の値を使用しなければなりません。
// 不一致の型をSwapするとパニックを引き起こし、Swap(nil)も同様です。
func (v *Value) Swap(new any) (old any)

// CompareAndSwapは、[Value] の比較交換操作を実行します。
//
// 与えられたValueに対するCompareAndSwapのすべての呼び出しは、同じ具体的な型の値を使用しなければなりません。
// 不一致の型をCompareAndSwapするとパニックを引き起こし、CompareAndSwap(old, nil)も同様です。
func (v *Value) CompareAndSwap(old, new any) (swapped bool)

// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect

import "github.com/shogo82148/std/iter"

// Seqは、vの要素をループするiter.Seq[Value]を返します。
// vの種類がFuncの場合、それは結果を持たず、型func(T) boolの単一の引数を取る
// 関数でなければなりません。
// vの種類がPointerの場合、ポインタ要素の型はArrayでなければなりません。
// それ以外の場合、vの種類はInt、Int8、Int16、Int32、Int64、
// Uint、Uint8、Uint16、Uint32、Uint64、Uintptr、
// Array、Chan、Map、Slice、またはStringでなければなりません。
func (v Value) Seq() iter.Seq[Value]

// Seq2は、vの要素をループするiter.Seq2[Value, Value]を返します。
// vの種類がFuncの場合、それは結果を持たず、
// 型func(K, V) boolの単一の引数を取る関数でなければなりません。
// vの種類がPointerの場合、ポインタ要素の型はArrayでなければなりません。
// それ以外の場合、vの種類はArray、Map、Slice、またはStringでなければなりません。
func (v Value) Seq2() iter.Seq2[Value, Value]

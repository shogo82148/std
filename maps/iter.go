// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maps

import "github.com/shogo82148/std/iter"

// Allは、mからキーと値のペアを反復するイテレータを返します。
// イテレーションの順序は指定されておらず、呼び出しごとに同じであることは保証されません。
func All[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V]

// Keysは、m内のキーを反復するイテレータを返します。
// イテレーションの順序は指定されておらず、呼び出しごとに同じであることは保証されません。
func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K]

// Valuesは、m内の値を反復するイテレータを返します。
// イテレーションの順序は指定されておらず、呼び出しごとに同じであることは保証されません。
func Values[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[V]

// Insertは、seqからmにキーと値のペアを追加します。
// seq内のキーがすでにmに存在する場合、その値は上書きされます。
func Insert[Map ~map[K]V, K comparable, V any](m Map, seq iter.Seq2[K, V])

// Collect collects key-value pairs from seq into a new map
// and returns it.
func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// mapsパッケージは、任意の型のマップに役立つさまざまな関数を定義します。
package maps

// Equalは、2つのマップが同じキー/値のペアを含むかどうかを報告します。
// 値は==を使用して比較されます。
func Equal[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool { return false }

// EqualFuncはEqualと同様ですが、eqを使用して値を比較します。
// キーは引き続き==で比較されます。
func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool {
	return false
}

// Cloneは、mのコピーを返します。これは浅いクローンです。
// 新しいキーと値は、通常の代入を使用して設定されます。
func Clone[M ~map[K]V, K comparable, V any](m M) M { return nil }

// Copyは、srcのすべてのキー/値ペアをコピーし、それらをdstに追加します。
// srcのキーがdstにすでに存在する場合、
// dstの値はsrcのキーに関連付けられた値によって上書きされます。
func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {}

// DeleteFuncは、delがtrueを返す場合、mから任意のキー/値ペアを削除します。
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool) {}

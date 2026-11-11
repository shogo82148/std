// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements commonly used type predicates.

package types2

// IsInterface reports whether t is an interface type.
func IsInterface(t Type) bool

// Comparable reports whether values of type T are comparable.
func Comparable(T Type) bool

// StrictlyComparable は、型 T の値が厳密に比較可能かどうかを報告します。
//
// 型が厳密に比較可能であるとは、その型が比較可能であり、
// 構造体のフィールドと配列の要素を再帰的にたどってもインターフェース型を含まないことをいいます。
// 正確な定義については https://go.dev/ref/spec#Comparison_operators を参照してください。
func StrictlyComparable(T Type) bool

// Default returns the default "typed" type for an "untyped" type;
// it returns the incoming type for all other types. The default type
// for untyped nil is untyped nil.
func Default(t Type) Type

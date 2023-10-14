// "go test -run=Generate -write=all"によって生成されたコードです。編集しないでください。

// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルではよく使用される型の述語を実装しています。

package types

// IsInterfaceはtがインターフェース型かどうかを報告します。
func IsInterface(t Type) bool

// Comparableは、型Tの値が比較可能かどうかを報告します。
func Comparable(T Type) bool

// Defaultは、「未指定の」型に対して「型付き」のデフォルト型を返します；
// 他のすべての型に対しては、入力された型を返します。未指定のnilのデフォルト型は未指定のnilです。
func Default(t Type) Type

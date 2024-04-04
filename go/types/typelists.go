<<<<<<< HEAD
// 「go test -run = Generate -write = all」によって生成されたコード；編集しないでください。
=======
// Code generated by "go test -run=Generate -write=all"; DO NOT EDIT.
// Source: ../../cmd/compile/internal/types2/typelists.go
>>>>>>> upstream/master

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// TypeParamListは型パラメータのリストを保持します。
type TypeParamList struct{ tparams []*TypeParam }

// Lenはリスト内の型パラメーターの数を返します。
// nilなレシーバーでも安全に呼び出すことができます。
func (l *TypeParamList) Len() int

// Atはリスト内のi番目の型パラメータを返します。
func (l *TypeParamList) At(i int) *TypeParam

// TypeListは型のリストを保持します。
type TypeList struct{ types []Type }

// Lenはリスト内の要素数を返します。
// nilの受信側で呼び出しても安全です。
func (l *TypeList) Len() int

// Atはリスト内のi番目のタイプを返します。
func (l *TypeList) At(i int) Type

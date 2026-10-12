// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import "github.com/shogo82148/std/iter"

// Methodsは、インターフェースのすべてのメソッドをId順に走査するgo1.23のイテレーターを返します。
//
// 例: for m := range t.Methods() { ... }
func (t *Interface) Methods() iter.Seq[*Func]

// ExplicitMethodsは、インターフェースの明示的なメソッドをId順に走査するgo1.23のイテレーターを返します。
//
// 例: for m := range t.ExplicitMethods() { ... }
func (t *Interface) ExplicitMethods() iter.Seq[*Func]

// EmbeddedTypesは、インターフェースに埋め込まれた型を走査するgo1.23のイテレーターを返します。
//
// 例: for e := range t.EmbeddedTypes() { ... }
func (t *Interface) EmbeddedTypes() iter.Seq[Type]

// Methodsは、名前付き型で宣言されたメソッドを走査するgo1.23のイテレーターを返します。
//
// 例: for m := range t.Methods() { ... }
func (t *Named) Methods() iter.Seq[*Func]

// Childrenは、スコープs内にネストされた子スコープを走査するgo1.23のイテレーターを返します。
//
// 例: for child := range scope.Children() { ... }
func (s *Scope) Children() iter.Seq[*Scope]

// Fieldsは、構造体型のフィールドを走査するgo1.23のイテレーターを返します。
//
// 例: for field := range s.Fields() { ... }
func (s *Struct) Fields() iter.Seq[*Var]

// Variablesは、タプル型の変数を走査するgo1.23のイテレーターを返します。
//
// 例: for v := range tuple.Variables() { ... }
func (t *Tuple) Variables() iter.Seq[*Var]

// Methodsは、メソッドセットのメソッドを走査するgo1.23のイテレーターを返します。
//
// 例: for method := range s.Methods() { ... }
func (s *MethodSet) Methods() iter.Seq[*Selection]

// Termsは、ユニオンの項を走査するgo1.23のイテレーターを返します。
//
// 例: for term := range union.Terms() { ... }
func (u *Union) Terms() iter.Seq[*Term]

// TypeParamsは、型パラメーターのリストを走査するgo1.23のイテレーターを返します。
//
// 例: for tparam := range l.TypeParams() { ... }
func (l *TypeParamList) TypeParams() iter.Seq[*TypeParam]

// Typesは、型のリストの要素を走査するgo1.23のイテレーターを返します。
//
// 例: for t := range l.Types() { ... }
func (l *TypeList) Types() iter.Seq[Type]

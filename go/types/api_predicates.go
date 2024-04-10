// Code generated by "go test -run=Generate -write=all"; DO NOT EDIT.
// Source: ../../cmd/compile/internal/types2/api_predicates.go

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements exported type predicates.

package types

// AssertableToは、型Vの値が型Tにアサートされることができるかどうかを報告します。
//
// AssertableToの動作は、3つのケースで未指定です：
//   - TがTyp[Invalid]である場合
//   - Vが一般化されたインタフェースである場合。つまり、Goコードで型制約としてのみ使用されるインタフェースである場合
//   - Tが未実体化のジェネリック型である場合
func AssertableTo(V *Interface, T Type) bool

// AssignableToは、型Vの値が型Tの変数に代入可能かどうかを報告します。
//
// AssignableToの動作は、VまたはTがTyp[Invalid]またはインスタンス化されていないジェネリック型の場合、指定されていません。
func AssignableTo(V, T Type) bool

// ConvertibleToは、型Vの値が型Tの値に変換可能かどうかを報告します。
//
// ConvertibleToの動作は、VまたはTがTyp[Invalid]またはインスタンス化されていないジェネリック型である場合、指定されていません。
func ConvertibleTo(V, T Type) bool

// Implementsは、型VがインターフェースTを実装しているかどうかを報告します。
//
// VがTyp[Invalid]やインスタンス化されていないジェネリック型の場合、Implementsの動作は未指定です。
func Implements(V Type, T *Interface) bool

// Satisfiesは型Vが制約Tを満たすかどうかを報告します。
//
// VがTyp[Invalid]またはインスタンス化されていないジェネリック型である場合、Satisfiesの動作は指定されていません。
func Satisfies(V Type, T *Interface) bool

// Identicalはxとyが同じ型であるかどうかを返します。
// [Signature] 型のレシーバは無視されます。
//
// [Identical]、[Implements]、[Satisfies] などの述語は、
// 両方のオペランドが一貫したシンボルのコレクション（[Object] 値）に属していると仮定します。
// 例えば、2つの [Named] 型は、それらの [Named.Obj] メソッドが同じ [TypeName] シンボルを返す場合にのみ同一となります。
// シンボルのコレクションが一貫しているとは、パスがPである各論理パッケージについて、
// それらのシンボルの作成には最大で一回の [NewPackage](P, ...)の呼び出しが関与していることを意味します。
// 一貫性を確保するために、すべてのロードされたパッケージとその依存関係に対して単一の [Importer] を使用します。
// 詳細は https://github.com/golang/go/issues/57497 を参照してください。
func Identical(x, y Type) bool

// IdenticalIgnoreTagsは、タグを無視した場合にxとyが同じ型であるかどうかを報告します。
// [Signature] 型のレシーバーは無視されます。
func IdenticalIgnoreTags(x, y Type) bool

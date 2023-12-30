// "go test -run=Generate -write=all" により生成されたコードです。編集しないでください。

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/go/constant"
	"github.com/shogo82148/std/go/token"
)

// Objectは、パッケージ、定数、型、変数、関数（メソッドを含む）、またはラベルなどの名前付きの言語エンティティを表します。
// すべてのオブジェクトはObjectインターフェースを実装しています。
type Object interface {
	Parent() *Scope
	Pos() token.Pos
	Pkg() *Package
	Name() string
	Type() Type
	Exported() bool
	Id() string

	String() string

	order() uint32

	color() color

	setType(Type)

	setOrder(uint32)

	setColor(color color)

	setParent(*Scope)

	sameId(pkg *Package, name string) bool

	scopePos() token.Pos

	setScopePos(pos token.Pos)
}

// Idが公開されていれば、そのままの名前を返します。それ以外の場合は、
// パッケージのパスで修飾された名前を返します。
func Id(pkg *Package, name string) string

// PkgNameはインポートされたGoパッケージを表します。
// PkgNameには型はありません。
type PkgName struct {
	object
	imported *Package
	used     bool
}

// NewPkgNameは、インポートされたパッケージを表す新しいPkgNameオブジェクトを返します。
// 残りの引数は、全てのオブジェクトで見つかった属性を設定します。
func NewPkgName(pos token.Pos, pkg *Package, name string, imported *Package) *PkgName

// Importedはインポートされたパッケージを返します。
// これはインポート文を含むパッケージとは異なります。
func (obj *PkgName) Imported() *Package

// Constは宣言された定数を表します。
type Const struct {
	object
	val constant.Value
}

// NewConstは値valを持つ新しい定数を返します。
// 残りの引数は、すべてのオブジェクトで見つかる属性を設定します。
func NewConst(pos token.Pos, pkg *Package, name string, typ Type, val constant.Value) *Const

// Valは定数の値を返します。
func (obj *Const) Val() constant.Value

// TypeNameは(定義済みまたはエイリアスの)型の名前を表します。
type TypeName struct {
	object
}

// NewTypeNameは、与えられたtypを指定する新しい型名を返します。
// 残りの引数は、すべてのオブジェクトで見つかった属性を設定します。
//
// typ引数は、定義済み（Named）タイプまたはエイリアスタイプである場合も可能です。
// また、nilである場合も、TypeNameを引数として使用でき、
// 副作用としてTypeNameのタイプがNewNamedに設定されます。
func NewTypeName(pos token.Pos, pkg *Package, name string, typ Type) *TypeName

// IsAliasは、objが型のエイリアス名であるかどうかを報告します。
func (obj *TypeName) IsAlias() bool

type Var struct {
	object
	embedded bool
	isField  bool
	used     bool
	origin   *Var
}

// NewVarは新しい変数を返します。
// 引数はすべてのオブジェクトで見つかった属性を設定します。
func NewVar(pos token.Pos, pkg *Package, name string, typ Type) *Var

// NewParam は関数のパラメータを表す新しい変数を返します。
func NewParam(pos token.Pos, pkg *Package, name string, typ Type) *Var

// NewFieldは、構造体のフィールドを表す新しい変数を返します。
// 埋め込まれたフィールドの場合、名前はフィールドがアクセス可能な
// 非修飾の型名です。
func NewField(pos token.Pos, pkg *Package, name string, typ Type, embedded bool) *Var

// Anonymousは変数が埋め込まれたフィールドかどうかを示します。
// Embeddedと同様ですが、後方互換性のために存在します。
func (obj *Var) Anonymous() bool

// Embeddedは変数が埋め込まれたフィールドかどうかを示します。
func (obj *Var) Embedded() bool

// IsFieldは、変数が構造体のフィールドであるかどうかを報告します。
func (obj *Var) IsField() bool

// Originは、そのレシーバのための正規のVar、つまりInfo.Defsに記録されたVarオブジェクトを返します。
//
// インスタンス化中に作成された合成Var（型引数に依存する構造体フィールドや
// 関数パラメータなど）については、これはジェネリック（インスタンス化されていない）型上の
// 対応するVarになります。他のすべてのVarについて、Originはレシーバを返します。
func (obj *Var) Origin() *Var

// Funcは、宣言された関数、具体的なメソッド、または抽象（インターフェース）メソッドを表します。そのType()は常に*Signatureです。
// 抽象メソッドは、埋め込みにより多くのインターフェースに所属することがあります。
type Func struct {
	object
	hasPtrRecv_ bool
	origin      *Func
}

// NewFuncは与えられたシグネチャを持つ新しい関数を返します。これは関数の型を表します。
func NewFunc(pos token.Pos, pkg *Package, name string, sig *Signature) *Func

// FullNameは関数またはメソッドobjのパッケージ名またはレシーバー型名で修飾された名前を返します。
func (obj *Func) FullName() string

// Scopeは関数の本体ブロックのスコープを返します。
// 結果は、インポートされたまたはインスタンス化された関数やメソッドに対してはnilです
// （ただし、インスタンス化された関数にアクセスするメカニズムもありません）。
func (obj *Func) Scope() *Scope

// Originは、レシーバーの正確なFunc、つまりInfo.Defsに記録されたFuncオブジェクトを返します。
// インスタンス化中に作成された合成関数（具名型のメソッドや型引数に依存するインターフェースのメソッドなど）の場合、これはジェネリック（インスタンス化されていない）型の対応するFuncになります。その他のすべてのFuncに対して、Originはレシーバーを返します。
func (obj *Func) Origin() *Func

// Pkg returns the package to which the function belongs.
//
// The result is nil for methods of types in the Universe scope,
// like method Error of the error built-in interface type.
func (obj *Func) Pkg() *Package

// Labelは宣言されたラベルを表します。
// ラベルはタイプを持ちません。
type Label struct {
	object
	used bool
}

// NewLabel は新しいラベルを返します。
func NewLabel(pos token.Pos, pkg *Package, name string) *Label

// Builtinは組み込み関数を表します。
// 組み込み関数には有効な型はありません。
type Builtin struct {
	object
	id builtinId
}

// Nilは、事前宣言された値であるnilを表します。
type Nil struct {
	object
}

// ObjectStringはobjの文字列形式を返します。
// Qualifierはパッケージレベルのオブジェクトの印刷を制御し、nilである可能性があります。
func ObjectString(obj Object, qf Qualifier) string

func (obj *PkgName) String() string
func (obj *Const) String() string
func (obj *TypeName) String() string
func (obj *Var) String() string
func (obj *Func) String() string
func (obj *Label) String() string
func (obj *Builtin) String() string
func (obj *Nil) String() string

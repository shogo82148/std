// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package reflect はランタイムリフレクションを実装し、プログラムが任意の型のオブジェクトを操作できるようにします。典型的な使用方法は、静的型 interface{} の値を取り、TypeOf を呼び出してその動的な型情報を抽出することです。TypeOf は Type を返します。
// ValueOf への呼び出しは、実行時データを表す Value を返します。Zero は Type を受け取り、その型のゼロ値を表す Value を返します。
// Go におけるリフレクションの紹介については、「The Laws of Reflection」を参照してください：
// https://golang.org/doc/articles/laws_of_reflection.html
package reflect

import "github.com/ethereum/go-ethereum/accounts/abi"

// TypeはGoの型の表現です。
//
// すべてのメソッドがすべての種類の型に適用されるわけではありません。制限がある場合は、各メソッドのドキュメントに記載されています。
// kind-specificメソッドを呼び出す前に型の種類を知るためにKindメソッドを使用してください。型の種類に適切でないメソッドを呼び出すと、ランタイムパニックが発生します。
//
// Typeの値は比較可能であり、==演算子などで使用することができます。
// 2つのTypeの値は、同一の型を表している場合に等しいとされます。
type Type interface {
	Align() int

	FieldAlign() int

	Method(int) Method

	MethodByName(string) (Method, bool)

	NumMethod() int

	Name() string

	PkgPath() string

	Size() uintptr

	String() string

	Kind() Kind

	Implements(u Type) bool

	AssignableTo(u Type) bool

	ConvertibleTo(u Type) bool

	Comparable() bool

	Bits() int

	ChanDir() ChanDir

	IsVariadic() bool

	Elem() Type

	Field(i int) StructField

	FieldByIndex(index []int) StructField

	FieldByName(name string) (StructField, bool)

	FieldByNameFunc(match func(string) bool) (StructField, bool)

	In(i int) Type

	Key() Type

	Len() int

	NumField() int

	NumIn() int

	NumOut() int

	Out(i int) Type

	common() *abi.Type
	uncommon() *uncommonType
}

// Kindは、 [Type] が表す特定の種類の型を表します。
// ゼロのKindは有効な種類ではありません。
type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)

// Ptrは [Pointer] 種別の旧名称です。
const Ptr = Pointer

// ChanDirはチャネルの方向を表します。
type ChanDir int

const (
	RecvDir ChanDir = 1 << iota
	SendDir
	BothDir = RecvDir | SendDir
)

// Methodは単一のメソッドを表します。
type Method struct {
	// Nameはメソッド名です。
	Name string

	// PkgPathは、小文字（エクスポートされていない）のメソッド名を修飾するパッケージパスです。大文字（エクスポートされた）のメソッド名の場合は空です。
	// PkgPathとNameの組み合わせは、メソッドセット内のメソッドを一意に識別します。
	// https://golang.org/ref/spec#Uniqueness_of_identifiers を参照してください。
	PkgPath string

	Type  Type
	Func  Value
	Index int
}

// IsExportedはメソッドがエクスポートされているかどうかを報告します。
func (m Method) IsExported() bool

// String は k の名前を返します。
func (k Kind) String() string

func (d ChanDir) String() string

// StructFieldはstruct内のフィールドを1つ記述します。
type StructField struct {
	// Nameはフィールド名です。
	Name string

	// PkgPathは小文字（エクスポートされていない）のフィールド名を修飾するパッケージパスです。大文字（エクスポートされた）のフィールド名には空です。
	// 詳細はhttps://golang.org/ref/spec#Uniqueness_of_identifiersを参照してください。
	PkgPath string

	Type      Type
	Tag       StructTag
	Offset    uintptr
	Index     []int
	Anonymous bool
}

// IsExportedはフィールドがエクスポートされているかどうかを報告します。
func (f StructField) IsExported() bool

// StructTagは、structフィールドのタグ文字列です。
//
// 慣例として、タグ文字列はオプションでスペースで区切られたkey:"value"の連結です。
// 各キーは、スペース（U+0020 ' '）、引用符（U+0022 '"'）、
// コロン（U+003A ':'）以外の非制御文字で構成される空でない文字列です。
// 各値は、U+0022 '"'文字とGoの文字列リテラル構文を使用して引用されます。
type StructTag string

// Getはtag文字列内のキーに関連付けられた値を返します。
// もしtag内にそのようなキーが存在しない場合、Getは空の文字列を返します。
// もしtagが従来の形式を持っていない場合、Getが返す値は不特定です。
// タグが明示的に空の文字列に設定されているかどうかを判断するには、Lookupを使用してください。
func (tag StructTag) Get(key string) string

// Lookupは、タグ文字列内のキーに関連する値を返します。
// キーがタグ内に存在する場合、その値（空かもしれません）が返されます。
// キーがタグに明示的に設定されていない場合、返される値は空の文字列になります。
// okの返り値は、値がタグ文字列に明示的に設定されているかどうかを報告します。
// タグに通常の形式がない場合、Lookupによって返される値は指定されていません。
func (tag StructTag) Lookup(key string) (value string, ok bool)

// TypeOfは、iの動的な型を表す反射 [Type] を返します。
// もしiがnilのインターフェース値である場合、TypeOfはnilを返します。
func TypeOf(i any) Type

// PtrToは、要素tを持つポインタ型を返します。
// 例えば、もしtがFoo型を表すなら、PtrTo(t)は*Fooを表します。
//
// PtrToは [PointerTo] の古い綴りです。
// これらの2つの関数は同じように動作します。
//
// Deprecated: [PointerTo] によって置き換えられました。
func PtrTo(t Type) Type

// PointerToは要素tを持つポインタ型を返します。
// 例えば、もしtがFoo型を表すなら、PointerTo(t)は*Fooを表します。
func PointerTo(t Type) Type

// ChanOfは指定された方向と要素の型を持つチャネル型を返します。
// たとえば、tがintを表す場合、ChanOf(RecvDir, t)は<-chan intを表します。
//
// gcのランタイムは、チャネルの要素型に64 kBの制限を課しています。
// もしtのサイズがこの制限以上である場合、ChanOfはパニックを発生させます。
func ChanOf(dir ChanDir, t Type) Type

// MapOfは与えられたキーと要素の型のマップタイプを返します。
// 例えば、もしkがintを表し、eがstringを表すならば、MapOf(k, e)はmap[int]stringを表します。
//
// もしキーの型が有効なマップキータイプではない場合（つまり、Goの==演算子を実装していない場合）、MapOfはパニックを起こします。
func MapOf(key, elem Type) Type

// FuncOfは与えられた引数と戻り値の型を持つ関数型を返します。
// 例えば、kがintを表し、eがstringを表す場合、
// FuncOf([]Type{k}, []Type{e}, false)はfunc(int) stringを表します。
//
// 可変引数の引数は、関数が可変引数かどうかを制御します。FuncOfは、
// variadicがtrueであり、in[len(in)-1]がスライスを表していない場合にパニックを起こします。
func FuncOf(in, out []Type, variadic bool) Type

// SliceOfは要素の型がtのスライス型を返します。
// 例えば、tがintを表す場合、SliceOf(t)は[]intを表します。
func SliceOf(t Type) Type

// StructOfはフィールドを含む構造体の型を返します。
// OffsetとIndexのフィールドは無視され、コンパイラによって計算されます。
//
// StructOfは、現在、埋め込みフィールドの昇格メソッドをサポートしておらず、
// エクスポートされていないStructFieldsが渡された場合にパニックを引き起こします。
func StructOf(fields []StructField) Type

// ArrayOfは、与えられた長さと要素の型を持つ配列型を返します。
// 例えば、tがintを表す場合、ArrayOf(5, t)は[5]intを表します。
//
// もし結果の型が利用可能なアドレススペースよりも大きくなる場合、ArrayOfはパニックを発生させます。
func ArrayOf(length int, elem Type) Type

// TypeFor returns the [Type] that represents the type argument T.
func TypeFor[T any]() Type
